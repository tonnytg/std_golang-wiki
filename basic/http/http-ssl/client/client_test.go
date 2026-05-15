package client

import (
	"testing"
)

func TestClient(t *testing.T) {
	statusCode, err := Client()
	if err != nil {
		t.Error(err)
	}
	if statusCode != 200 {
		t.Errorf("status expected 200 but got %d", statusCode)
	}
}
