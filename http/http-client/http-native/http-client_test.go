package http_client_test

import (
	httpClient "github.com/tonnytg/std-golang-wiki/http-client"
	"testing"
)

func TestGetSite(t *testing.T) {
	url := "https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist"
	err := httpClient.GetSite(url)
	if err != nil {
		t.Errorf("can't use get: %v", err)
	}
}

func TestPostSite(t *testing.T) {
	url := "https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist"
	err := httpClient.PostSite(url)
	if err != nil {
		t.Errorf("cant use post: %v", err)
	}
}