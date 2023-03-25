package userinfo

type Details []Detail

type Detail struct {
	FirstName string  `json:"firstName,omitempty" db:"first_name"`
	LastName  string  `json:"lastName,omitempty" db:"last_name"`
	Email     string  `json:"email,omitempty"  db:"email"`
	Username  string  `json:"username,omitempty"  db:"username"`
	Password  string  `json:"password,omitempty"  db:"password"`
	Address   Address `json:"address,omitempty" db:"address"`
}
