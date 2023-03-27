package userinfo

type Emails []Email

type Email struct {
	Address   string `json:"address,omitempty" db:"address"`
	IsPrimary bool   `json:"isPrimary,omitempty" db:"is_primary"`
}
