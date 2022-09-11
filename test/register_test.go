package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erviangelar/go-user-api/models"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	hTest := Database()
	// err := refreshUserTable()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	user, err := hTest.seedOneUser()
	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}

	samples := []struct {
		username     string
		password     string
		errorMessage string
	}{
		{
			username:     user.Username,
			password:     "password1234",
			errorMessage: "",
		},
		{
			username:     user.Username,
			password:     "Wrong password",
			errorMessage: "crypto/bcrypt: hashedPassword is not the hash of the given password",
		},
		{
			username:     "Wrong email",
			password:     "password",
			errorMessage: "record not found",
		},
	}

	for _, v := range samples {
		request := models.RegisterRequest{Username: v.username, Password: v.password, ConfirmPassword: v.password}
		user, err := hTest.handler.Create(&request)
		if err != nil {
			assert.Equal(t, err, errors.New(v.errorMessage))
		} else {
			assert.NotNil(t, user)
		}
	}
}

func TestRegisters(t *testing.T) {

	hTest := Database()
	hTest.refreshUserTable()

	_, err := hTest.seedOneUser()
	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		username     string
		password     string
		errorMessage string
	}{
		{
			inputJSON:    `{"username": "pet@gmail.com", "password": "password"}`,
			statusCode:   200,
			errorMessage: "",
		},
		{
			inputJSON:    `{"username": "pet@gmail.com", "password": "wrong password"}`,
			statusCode:   422,
			errorMessage: "Incorrect Password",
		},
		{
			inputJSON:    `{"username": "frank@gmail.com", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Incorrect Details",
		},
		{
			inputJSON:    `{"username": "kangmail.com", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Invalid Email",
		},
		{
			inputJSON:    `{"username": "", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Required Email",
		},
		{
			inputJSON:    `{"username": "kan@gmail.com", "password": ""}`,
			statusCode:   422,
			errorMessage: "Required Password",
		},
		{
			inputJSON:    `{"username": "", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Required Email",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/register", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		hTest.handler.Router.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 200 {
			assert.NotEqual(t, rr.Body.String(), "")
		}

		if v.statusCode == 422 && v.errorMessage != "" {
			responseMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to json: %v", err)
			}
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}
