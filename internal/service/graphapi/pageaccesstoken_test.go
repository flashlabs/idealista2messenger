package graphapi_test

import (
	"reflect"
	"testing"

	"github.com/flashlabs/idealista2messenger/internal/service/graphapi"

	_ "github.com/flashlabs/idealista2messenger/internal/test"
)

func TestPageAccessTokenFromFile(t *testing.T) {
	type args struct {
		file string
	}

	tests := []struct {
		name    string
		args    args
		want    *graphapi.PageAccessToken
		wantErr bool
	}{
		{
			name: "successfully initialized",
			args: args{
				file: "config/page_access_token.json.dist",
			},
			want:    &graphapi.PageAccessToken{Token: "<PAGE ACCESS TOKEN GOES HERE>"},
			wantErr: false,
		},
		{
			name: "invalid accecc token path",
			args: args{
				file: "invalid_patj/page_access_token.json.dist",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := graphapi.PageAccessTokenFromFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("PageAccessTokenFromFile() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PageAccessTokenFromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
