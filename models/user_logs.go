package models

type UserLogs struct {
	BaseMode
	Activity  string `json:"activity"`
	IpRequest string `json:"ip-request"`
	UserAgent string `json:"user-agent"`
}
