package models

type User struct {
	Name        string `json:"name"`
	Email       string `json:"emailAddress"`
	DisplayName string `json:"displayName"`
}
