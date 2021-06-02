package repository

import (
	"reflect"
	"testing"

	usermodel "github.com/mreines/go-users-api/model"
)

func TestUsersGet(t *testing.T) {
	tests := []struct {
		name string
		want []usermodel.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UsersGet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersPost(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    usermodel.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UsersPost(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersPost() = %v, want %v", got, tt.want)
			}
		})
	}
}
