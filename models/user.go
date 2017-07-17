package models

type User struct {
	Id                 int    `json:"id"`
	Email              string `json:"email"`
	Password           string `json:"password,omitempty"`
	PasswordHash       string `json:"-"`
	PasswordResetToken string `json:"-"`
	Address            string `json:"address"`
	Provider           string `json:"provider,omitempty"`
	ProviderId         string `json:"provider_id,omitempty"`
}
