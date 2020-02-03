package crypto

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/edznux/wonderxss/api"
)

func TestGetJWTToken(t *testing.T) {
	type args struct {
		user api.User
		key  string
	}
	tests := []struct {
		name    string
		args    args
		isOk    func(input string) bool
		wantErr bool
	}{
		{
			name: "Test JWT Creation",
			args: args{
				user: api.User{
					ID:               "1234-4567-8901",
					Username:         "test",
					TwoFactorEnabled: false,
					CreatedAt:        time.Now(),
					ModifiedAt:       time.Now(),
				},
				key: "fakekey",
			},
			isOk: func(input string) bool {
				splitted := strings.Split(input, ".")
				if len(splitted) != 3 {
					return false
				}
				if splitted[0] == "" || splitted[1] == "" || splitted[2] == "" {
					return false
				}
				// TODO: We should add base64 decoding and stuff in tests
				return true
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetJWTToken(tt.args.user, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJWTToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.isOk(got) {
				t.Errorf("GetJWTToken() = %v", got)
			}
		})
	}
}

func TestVerifyJWTToken(t *testing.T) {
	type args struct {
		tokenString string
		key         string
	}
	tests := []struct {
		name    string
		args    args
		want    jwt.Claims
		wantErr bool
	}{
		{
			name: "Verify empty token with standard key",
			args: args{
				key:         "some-fixed-test-key",
				tokenString: "",
			},
			wantErr: true,
		},
		{
			name: "Verify malformated/invalid base64 token with standard key",
			args: args{
				key:         "some-fixed-test-key",
				tokenString: "lol.aze.aze",
			},
			wantErr: true,
		},
		{
			name: "Verify expired token with standard key",
			args: args{
				key:         "some-fixed-test-key",
				tokenString: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoiYm9iIiwidXNlcl9uYW1lIjoiYm9iIiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNTgwNjE0NTA0LCJqdGkiOiJhZTI5NGQ1YS03ZTIzLTRjMjctOGFlNy02ZDdjMDk3MGI5YmIiLCJpYXQiOjE1ODA2MTA5MDR9.N3LjjBl7mvb9GwDKmTJWnB8goXE1c3IbUTsnXp7RZ4w",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyJWTToken(tt.args.tokenString, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyJWTToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyJWTToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
