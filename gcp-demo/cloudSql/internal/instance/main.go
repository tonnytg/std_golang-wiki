package instance

import (
	"encoding/json"
	"fmt"
	"local-p/internal/web"
	"time"
)

type Instance struct {
	Kind            string `json:"kind"`
	State           string `json:"state"`
	DatabaseVersion string `json:"databaseVersion"`
	Settings        struct {
		AuthorizedGaeApplications []interface{} `json:"authorizedGaeApplications"`
		Tier                      string        `json:"tier"`
		Kind                      string        `json:"kind"`
		AvailabilityType          string        `json:"availabilityType"`
		PricingPlan               string        `json:"pricingPlan"`
		ReplicationType           string        `json:"replicationType"`
		ActivationPolicy          string        `json:"activationPolicy"`
		IpConfiguration           struct {
			AuthorizedNetworks []interface{} `json:"authorizedNetworks"`
			Ipv4Enabled        bool          `json:"ipv4Enabled"`
		} `json:"ipConfiguration"`
		LocationPreference struct {
			Zone string `json:"zone"`
			Kind string `json:"kind"`
		} `json:"locationPreference"`
		DataDiskType        string `json:"dataDiskType"`
		BackupConfiguration struct {
			StartTime               string `json:"startTime"`
			Kind                    string `json:"kind"`
			BackupRetentionSettings struct {
				RetentionUnit   string `json:"retentionUnit"`
				RetainedBackups int    `json:"retainedBackups"`
			} `json:"backupRetentionSettings"`
			Enabled                     bool `json:"enabled"`
			TransactionLogRetentionDays int  `json:"transactionLogRetentionDays"`
		} `json:"backupConfiguration"`
		ConnectorEnforcement   string `json:"connectorEnforcement"`
		SettingsVersion        string `json:"settingsVersion"`
		StorageAutoResizeLimit string `json:"storageAutoResizeLimit"`
		StorageAutoResize      bool   `json:"storageAutoResize"`
		DataDiskSizeGb         string `json:"dataDiskSizeGb"`
	} `json:"settings"`
	Etag        string `json:"etag"`
	IpAddresses []struct {
		Type      string `json:"type"`
		IpAddress string `json:"ipAddress"`
	} `json:"ipAddresses"`
	ServerCaCert struct {
		Kind             string    `json:"kind"`
		CertSerialNumber string    `json:"certSerialNumber"`
		Cert             string    `json:"cert"`
		CommonName       string    `json:"commonName"`
		Sha1Fingerprint  string    `json:"sha1Fingerprint"`
		Instance         string    `json:"instance"`
		CreateTime       time.Time `json:"createTime"`
		ExpirationTime   time.Time `json:"expirationTime"`
	} `json:"serverCaCert"`
	InstanceType               string    `json:"instanceType"`
	Project                    string    `json:"project"`
	ServiceAccountEmailAddress string    `json:"serviceAccountEmailAddress"`
	BackendType                string    `json:"backendType"`
	SelfLink                   string    `json:"selfLink"`
	ConnectionName             string    `json:"connectionName"`
	Name                       string    `json:"name"`
	Region                     string    `json:"region"`
	GceZone                    string    `json:"gceZone"`
	DatabaseInstalledVersion   string    `json:"databaseInstalledVersion"`
	MaintenanceVersion         string    `json:"maintenanceVersion"`
	CreateTime                 time.Time `json:"createTime"`
}

func Create(project string) {

	url := fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances", project)

	var data []string

	dataJson := `{
			"name": "my-instance",
			"rootPassword": "teste!123",
			"settings": {
				"tier": "db-n1-standard-1"
			}
		}`
	data = append(data, dataJson)

	body, err := web.Core("POST", url, 20, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)
}

func Get(project string, instance string) Instance {

	url := fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s", project, instance)

	body, err := web.Core("GET", url, 20, nil)
	if err != nil {
		panic(err)
	}
	var i Instance
	json.Unmarshal(body, &i)
	//fmt.Printf("%s\n", i)
	//fmt.Printf("Status: %s\n", i.State)
	return i
}

func Update(project string, instance string) {
	//https://sqladmin.googleapis.com/sql/v1beta4/projects/{project}/instances/{instance}

	url := fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s", project, instance)

	var data []string

	dataJson := `{
			"settings": {
				"ipConfiguration": {
					"authorizedNetworks": [
						{
							"name": "CasaGenerico",
							"value": "189.50.222.0/24"
						}
					]
				}
			}
		}`
	data = append(data, dataJson)

	body, err := web.Core("PATCH", url, 20, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", body)

}

func Delete(project string, instance string) {

	url := fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s", project, instance)

	body, err := web.Core("DELETE", url, 20, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", body)
}
