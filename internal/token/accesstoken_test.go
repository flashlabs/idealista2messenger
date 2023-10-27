package token_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"

	"github.com/flashlabs/idealista2messenger/internal/token"

	_ "github.com/flashlabs/idealista2messenger/internal/test"
)

func TestAccessTokenFromFile(t *testing.T) {
	type args struct {
		file string
	}

	tests := []struct {
		name    string
		args    args
		want    *oauth2.Token
		wantErr bool
	}{
		{
			name: "successfully read token",
			args: args{file: "config/token.json.dist"},
			want: &oauth2.Token{
				AccessToken:  "access_token",
				TokenType:    "Bearer",
				RefreshToken: "refresh_token",
			},
			wantErr: false,
		},
		{
			name:    "invalid token path",
			args:    args{file: "invalid_path/token.json.dist"},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := token.AccessTokenFromFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessTokenFromFile() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessTokenFromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveAccessToken(t *testing.T) {
	path := "/tmp/i2m_token_path"

	// remove if previously existed
	_ = os.Remove(path)

	assert.NoFileExists(t, path)
	token.SaveAccessToken(path, &oauth2.Token{
		AccessToken:  "access_token",
		TokenType:    "Bearer",
		RefreshToken: "refresh_token",
	})

	assert.FileExists(t, path)

	// cleanup
	_ = os.Remove(path)
}
