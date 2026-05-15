package project

import (
	"fmt"
	"local-p/internal/web"
)

func List() {

	// GCP List projects
	url := "https://cloudresourcemanager.googleapis.com/v1/projects"

	body, err := web.Core("GET", url, 20, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
