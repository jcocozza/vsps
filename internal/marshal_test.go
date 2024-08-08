package internal

import (
	"testing"

)

func TestMarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   Account
		wantErr bool
		want    []byte
	}{
		{"test1", Account{Name: "my account", Username: "user", Password: "pass"}, false, []byte("my account:\n    username: user\n    password: pass\n")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, didn't get one")
				}
			}
			if string(got) != string(tt.want) {
				t.Errorf("expected %s, got: %s", tt.want, got)
			}
		})
	}
}
