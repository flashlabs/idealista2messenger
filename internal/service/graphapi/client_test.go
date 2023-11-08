package graphapi_test

import (
	"reflect"
	"testing"

	"github.com/flashlabs/idealista2messenger/internal/service/graphapi"
	_ "github.com/flashlabs/idealista2messenger/internal/test"
)

func TestNewGraphApiClient(t *testing.T) {
	type args struct {
		PageAccessTokenFileLocation string
		PageId                      string
	}

	tests := []struct {
		name    string
		args    args
		want    graphapi.Client
		wantErr bool
	}{
		{
			name: "initialize client",
			args: args{
				PageAccessTokenFileLocation: "config/page_access_token.json.dist",
				PageId:                      "page_id",
			},
			want: graphapi.Client{
				PageAccessToken: &graphapi.PageAccessToken{Token: "<PAGE ACCESS TOKEN GOES HERE>"},
				PageID:          "page_id",
			},
			wantErr: false,
		},
		{
			name: "invalid client initialization",
			args: args{
				PageAccessTokenFileLocation: "invalid_token/page_access_token.json.dist",
				PageId:                      "page_id",
			},
			want:    graphapi.Client{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := graphapi.NewGraphApiClient(tt.args.PageAccessTokenFileLocation, tt.args.PageId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGraphApiClient() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraphApiClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
