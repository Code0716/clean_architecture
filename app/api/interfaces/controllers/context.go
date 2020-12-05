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
