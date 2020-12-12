package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	// MultipartForm()
	// PostForm()
	// SaveUploadedFile(interface{}, string)
	JSON(int, interface{})
}

type Response struct {
	Status       int         `json:"status"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

// Meta struct : response data
type Meta struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message"`
}

type AuthResponse struct {
	Status        int    `json:"status"`
	ErrorMessage  string `json:"error_message"`
	Authorization string `json:"Authorization"`
}
