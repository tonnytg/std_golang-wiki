package database

import (
	"fmt"
	"local-p/internal/web"
)

func Create(project string, instance string) {

	// https://sqladmin.googleapis.com/sql/v1beta4/projects/{project}/instances/{instance}/databases
	url := fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s/databases",
		project, instance)

	var data []string

	dataJson := `{
			"instance": "my-instance",
			"name": "mysql-db",
			"project": "totemic-chalice-364501"
		}`
	data = append(data, dataJson)

	body, err := web.Core("POST", url, 20, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)
}

func Get(project string, instance string, database string) {
	//https://sqladmin.googleapis.com/sql/v1beta4/projects/{project}/instances/{instance}/databases/{database}

	url := fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s/databases/%s", project, instance, database)

	body, err := web.Core("GET", url, 20, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)
}

func AddUser(project string, instance string) {
	// https://sqladmin.googleapis.com/sql/v1beta4/projects/{project}/instances/{instance}/users
	url := fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s/users", project, instance)

	var data []string

	dataJson := `{
			"instance": "my-instance",
			"name": "teste",
			"password": "teste!123",
			"project": "totemic-chalice-364501"
		}`
	data = append(data, dataJson)

	body, err := web.Core("POST", url, 20, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)
}
