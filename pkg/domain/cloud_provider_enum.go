// Package domain provides the domain model for the application.
package domain

const (
	// Amazon is the ID for Amazon cloud provider
	Amazon = "8f391035-f7fa-459d-8715-6a65c52729ae"
	// GCP is the ID for GCP cloud provider
	GCP = "ee5c73ba-c6a3-4cb4-888c-1c78b342a349"
	// Azure is the ID for Azure cloud provider
	Azure = "4b48a974-c2b5-4560-8218-f06aba7ce91b"
)

// CloudProvider represents a cloud provider
type CloudProvider struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var cloudProviders = []CloudProvider{
	{
		ID:   Amazon,
		Name: "Amazon",
	},
	{
		ID:   GCP,
		Name: "Google Cloud Platform",
	},
	{
		ID:   Azure,
		Name: "Microsoft Azure",
	},
}

// GetCloudProviders returns a list of cloud providers
func GetCloudProviders() []CloudProvider {
	return cloudProviders
}
