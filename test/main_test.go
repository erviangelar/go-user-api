package test

import (
	"database/sql"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/erviangelar/go-user-api/handler"
	"github.com/erviangelar/go-user-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type hTest struct {
	Mock    sqlmock.Sqlmock
	handler *handler.Handler
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	code := m.Run()
	os.Exit(code)
}

func Database() *hTest {
	var handlerTest = handler.Handler{}
	var (
		db   *sql.DB
		err  error
		mock sqlmock.Sqlmock
	)
	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("[sqlmock new] %s", err)
	}
	// defer db.Close()
	dialector := postgres.New(postgres.Config{
		Conn: db,
	})

	// open the database
	log.Println("Test")
	handlerTest.DB, err = gorm.Open(dialector, &gorm.Config{})
	// configs := config.LoadConfig()
	// db, err := gorm.Open(postgres.Open(config.DBConn), &gorm.Config{})
	log.Println(err)
	if err != nil {
		log.Println(err)
		log.Fatalf("[gorm open] %s", err)
	}
	handlerTest.DB.Logger.LogMode(logger.Info)
	return &hTest{
		handler: &handlerTest,
		Mock:    mock,
	}
}

func (h hTest) refreshUserTable() error {
	log.Println("migrate")
	err := h.handler.DB.AutoMigrate(&models.User{})
	log.Fatalf("[gorm migrate] %s", err)
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table(s)")
	return nil
}

func (h hTest) seedOneUser() (models.User, error) {

	// err := h.refreshUserTable()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	user := models.User{
		Name:     "Pet",
		Username: "pet@gmail.com",
		Password: "password1234",
	}
	h.Mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("id","username", "password") 
			VALUES ($1,$2,$3) RETURNING "users"."id"`)).
		WithArgs(user.Name, user.Username, user.Password).
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).AddRow(user.Name))

	log.Println("seed")
	err := h.handler.DB.Create(&user).Error
	if err != nil {
		log.Fatalf("Failed Seed %s", err)
		return models.User{}, err
	}
	return user, nil
}
