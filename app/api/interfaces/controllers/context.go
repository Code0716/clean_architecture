package controllers

type Context interface {
	// MultipartForm()
	// PostForm()
	// SaveUploadedFile(interface{}, string)
	JSON(int, interface{})
}
