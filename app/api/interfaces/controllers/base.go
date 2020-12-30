package controllers

import (
	"net/http"

	"github.com/Code0716/clean_architecture/app/api/domain"
)

type Base struct {
}

func (b *Base) FormatResponse(status int, data ...interface{}) *domain.Response {
	response := new(domain.Response)
	switch status {
	case http.StatusOK:
		meta := &domain.Meta{
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
		meta := &domain.Meta{
			Status: status,
		}
		response.Meta = meta
	}

	return response
}

func (b *Base) FormatAuthResponse(status int, token string) *domain.AuthResponse {
	response := new(domain.AuthResponse)
	switch status {
	case http.StatusOK:
		meta := &domain.Meta{
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
		meta := &domain.Meta{
			Status: status,
		}
		response.Meta = meta
	}

	return response
}
