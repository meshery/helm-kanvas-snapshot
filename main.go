package main

import (
	cmd "github.com/meshery/helm-kanvas-snapshot/cmd/kanvas-snapshot"
)

var (
	providerToken          string
	mesheryCloudAPIBaseURL string
	mesheryAPIBaseURL      string
	workflowAccessToken    string
)

func main() {
	cmd.Main(providerToken, mesheryCloudAPIBaseURL, mesheryAPIBaseURL, workflowAccessToken)
}
