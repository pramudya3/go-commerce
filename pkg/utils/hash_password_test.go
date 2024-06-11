package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		pass []byte
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successfully",
			args: args{
				pass: []byte("test"),
			},
			wantErr: false,
		},
		{
			name: "empty password",
			args: args{
				pass: []byte(""),
			},
			wantErr: false,
		},
		{
			name: "password too long",
			args: args{
				pass: []byte("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"), // contoh panjang lebih dari 72
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got == "" != tt.wantErr {
				t.Errorf("HashPassword() failed hashing password, got %v", got)
			}
		})
	}

}
