package models

type Author struct {
	Self        string `json:"self"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
}
