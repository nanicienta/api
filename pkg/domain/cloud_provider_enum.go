package domain

const (
	AMAZON = "8f391035-f7fa-459d-8715-6a65c52729ae"
	GCP    = "ee5c73ba-c6a3-4cb4-888c-1c78b342a349"
	AZURE  = "4b48a974-c2b5-4560-8218-f06aba7ce91b"
)

type CloudProvider struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var cloudProviders = []CloudProvider{
	{
		ID:   AMAZON,
		Name: "Amazon",
	},
	{
		ID:   GCP,
		Name: "Google Cloud Platform",
	},
	{
		ID:   AZURE,
		Name: "Microsoft Azure",
	},
}

func GetCloudProviders() []CloudProvider {
	return cloudProviders
}
