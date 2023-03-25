package models

import (
	"github.com/calebtraceyco/mind-your-business-api/external/models/userinfo"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Users []User

type User struct {
	ID     pgtype.UUID     `json:"id,omitempty" db:"id"`
	Detail userinfo.Detail `json:"detail,omitempty" db:"detail"`

	Emails    userinfo.Emails    `json:"emails,omitempty" db:"emails"`
	Addresses userinfo.Addresses `json:"addresses,omitempty" db:"addresses"`
	Contacts  userinfo.Contacts  `json:"contacts,omitempty" db:"contacts"`

	Token        string    `json:"token,omitempty"  db:"token"`
	RefreshToken string    `json:"refreshToken,omitempty"  db:"refresh_Token"`
	CreatedAt    time.Time `json:"createdAt"  db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt"  db:"updated_at"`
}
