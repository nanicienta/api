package command

type RefreshTokenCommand struct {
	RefreshToken string `json:"refreshToken"`
	DeviceID     string
}
