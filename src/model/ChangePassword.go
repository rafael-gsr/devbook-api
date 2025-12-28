package model

// ChangePassword is the struct that should be used as DTO for changeing the user password
type ChangePassword struct {
	New string `json:"new"`
	Old string `json:"old"`
}
