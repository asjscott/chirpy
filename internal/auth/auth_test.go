package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"ValidPassword", "mysecretpassword", false},
		{"EmptyPassword", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(hash) == 0 {
				t.Errorf("expected hash to be non-empty")
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{"ValidPassword", "mysecretpassword", func() string { h, _ := HashPassword("mysecretpassword"); return h }(), false},
		{"WrongPassword", "wrongpassword", func() string { h, _ := HashPassword("mysecretpassword"); return h }(), true},
		{"EmptyPassword", "", func() string { h, _ := HashPassword("mysecretpassword"); return h }(), true},
		{"EmptyHash", "mysecretpassword", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match, err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) || match != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}