// Package command provides the command structure for different actions on the service.
package command

// UpdateProfileCommand represents the command to update a user's profile.
type UpdateProfileCommand struct {
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Name       string `json:"name"`
	PhotoURL   string `json:"photoUrl"`
	Timezone   string `json:"timezone"`
	MuteSounds bool   `json:"muteSounds"`
}
