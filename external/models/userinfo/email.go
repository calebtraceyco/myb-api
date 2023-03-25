package userinfo

type Emails []Email

type Email struct {
	Address   string
	isPrimary bool
}
