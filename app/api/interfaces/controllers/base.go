package controllers

import (
	"net/http"
)

type Base struct {
}

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

func (b *Base) FormatResponse(status int, data ...interface{}) *Response {
	response := new(Response)
	switch status {
	case http.StatusOK:
		meta := &Meta{
			Status: http.StatusOK,
		}
		response.Meta = meta
		response.Data = data
	case http.StatusBadRequest: //400
		fallthrough
	case http.StatusUnauthorized: //401
		fallthrough
	case http.StatusForbidden: //403
		fallthrough
	case http.StatusNotFound: //404
		fallthrough
	case http.StatusConflict: //409
		fallthrough
	case http.StatusInternalServerError: //500
		meta := &Meta{
			Status: status,
		}
		response.Meta = meta
	}

	return response
}

func (b *Base) FormatAuthResponse(status int, token string) *AuthResponse {
	response := new(AuthResponse)
	switch status {
	case http.StatusOK:
		meta := &Meta{
			Status: http.StatusOK,
		}
		response.Meta = meta
		response.Authorization = token
	case http.StatusBadRequest: //400
		fallthrough
	case http.StatusUnauthorized: //401
		fallthrough
	case http.StatusForbidden: //403
		fallthrough
	case http.StatusNotFound: //404
		fallthrough
	case http.StatusConflict: //409
		fallthrough
	case http.StatusInternalServerError: //500
		meta := &Meta{
			Status: status,
		}
		response.Meta = meta
	}

	return response
}
