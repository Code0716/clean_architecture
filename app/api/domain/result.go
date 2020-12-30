package domain

type Response struct {
	Meta *Meta       `json:"meta"`
	Data interface{} `json:"data"`
}

type AuthResponse struct {
	Meta          *Meta  `json:"meta"`
	Authorization string `json:"Authorization"`
}

// Meta struct : response data
type Meta struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message"`
}
