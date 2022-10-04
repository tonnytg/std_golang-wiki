#!/bin/bash

gcloud auth application-default login

export GCP_TOKEN=`gcloud auth application-default print-access-token`

go fmt ./...

go run main.go