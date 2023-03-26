package psql

import (
	"github.com/calebtraceyco/mind-your-business-api/external"
	"github.com/calebtraceyco/mind-your-business-api/external/models"
	"github.com/calebtraceyco/mind-your-business-api/external/models/userinfo"
	"github.com/calebtraceyco/mind-your-business-api/internal/dao/psql"
	"github.com/jackc/pgx/v5/pgtype"
	"testing"
	"time"
)

type payload struct {
	Request  external.Request
	Endpoint string
}

var mockPayload = payload{Request: external.Request{
	User: &models.User{
		ID: pgtype.UUID{},
		Detail: userinfo.Detail{
			FirstName: "TEST",
			LastName:  "TEST",
			Email:     "TEST",
			Username:  "TEST",
			Password:  "TEST",
			Address:   userinfo.Address{},
		},
		Emails:       userinfo.Emails{},
		Addresses:    userinfo.Addresses{},
		Contacts:     userinfo.Contacts{},
		Token:        "TEST",
		RefreshToken: "TEST",
		CreatedAt:    mockTime,
		UpdatedAt:    mockTime,
	},
},
}

func TestMapper_PostgresExec(t *testing.T) {
	tests := []struct {
		name    string
		request *external.ApiRequest
		want    string
	}{
		{
			name: "Happy Path",
			request: &external.ApiRequest{
				Payload: struct {
					Request  external.Request `json:"request,omitempty"`
					Endpoint string           `json:"endpoint,omitempty"`
				}(mockPayload),
			},
			want: mockExec,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := psql.Mapper{}
			if got := m.NewUserExec(tt.request); got != tt.want {
				t.Errorf("NewUserExec() = %v, want %v", got, tt.want)
			}
		})
	}
}

var mockTime = getFakeTime()

func getFakeTime() time.Time {
	year, month, day := 2023, time.March, 22
	hour, min, sec := 10, 30, 0
	return time.Date(year, month, day, hour, min, sec, 0, time.UTC)
}

const mockExec = `insert into users (first_name, last_name, email, username, password, token, refresh_Token) values ('TEST', 'TEST', 'TEST', 'TEST', 'TEST', 'TEST', 'TEST')`
