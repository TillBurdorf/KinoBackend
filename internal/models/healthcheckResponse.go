package models

type HealthcheckResponse struct {
	Status      string
	Environment string
	Version     string
}
