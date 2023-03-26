package user

import (
	"github.com/calebtraceyco/mind-your-business-api/external"
	"github.com/calebtraceyco/mind-your-business-api/external/models"
	"github.com/calebtraceyco/mind-your-business-api/external/models/userinfo"
	"github.com/jackc/pgx/v5/pgtype"
	"testing"
	"time"
)

var mockTime = getFakeTime()

func getFakeTime() time.Time {
	year, month, day := 2023, time.March, 22
	hour, min, sec := 10, 30, 0
	return time.Date(year, month, day, hour, min, sec, 0, time.UTC)
}

var mockUser = external.Request{
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
}

func Test_mapFields(t *testing.T) {
	type args struct {
		request  any
		daoModel any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Path",
			args: args{
				request:  mockUser,
				daoModel: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mapFields(tt.args.request, tt.args.daoModel); (err != nil) != tt.wantErr {
				t.Errorf("mapFields() error = %v, wantErr %v", err, tt.wantErr)
			}
			//assert.Equal(t)
		})
	}
}
