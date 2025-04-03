package domain

type ErrorTypeBody struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

type ErrorType struct {
	Body       ErrorTypeBody
	StatusCode int
}
