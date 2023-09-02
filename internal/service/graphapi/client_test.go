package graphapi

import (
	"reflect"
	"testing"

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
		want    Client
		wantErr bool
	}{
		{
			name: "initialize client",
			args: args{
				PageAccessTokenFileLocation: "config/page_access_token.json.dist",
				PageId:                      "page_id",
			},
			want: Client{
				PageAccessToken: &PageAccessToken{Token: "<PAGE ACCESS TOKEN GOES HERE>"},
				PageId:          "page_id",
			},
			wantErr: false,
		},
		{
			name: "invalid client initialization",
			args: args{
				PageAccessTokenFileLocation: "invalid_token/page_access_token.json.dist",
				PageId:                      "page_id",
			},
			want:    Client{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGraphApiClient(tt.args.PageAccessTokenFileLocation, tt.args.PageId)
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
