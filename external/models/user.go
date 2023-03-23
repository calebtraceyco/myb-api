package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type User struct {
	ID           pgtype.UUID `json:"id,omitempty" db:"id"`
	FirstName    string      `json:"firstName,omitempty" db:"first_name"`
	LastName     string      `json:"lastName,omitempty" db:"last_name"`
	Email        string      `json:"email,omitempty"  db:"email"`
	Username     string      `json:"username,omitempty"  db:"username"`
	Password     string      `json:"password,omitempty"  db:"password"`
	Token        string      `json:"token,omitempty"  db:"token"`
	RefreshToken string      `json:"refreshToken,omitempty"  db:"refresh_Token"`
	CreatedAt    time.Time   `json:"createdAt"  db:"created_at"`
	UpdatedAt    time.Time   `json:"updatedAt"  db:"updated_at"`
}
