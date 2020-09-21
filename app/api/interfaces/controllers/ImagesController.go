package controllers

import (
	"../../usecase"
	"../database"
)

// images

type ImagesController struct {
	Interactor usecase.ImagesInteractor
}

func NewImagesController(sqlHandler database.SqlHandler) *ImagesController {
	return &ImagesController{
		Interactor: usecase.ImagesInteractor{
			ImagesRepository: &database.ImagesRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ImagesController) GetAll(c Context) {
	images, err := controller.Interactor.GetAllImages()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, images)
}

// Pre images
type PreImagesController struct {
	Interactor usecase.PreImagesInteractor
}

func NewPreImagesController(sqlHandler database.SqlHandler) *PreImagesController {
	return &PreImagesController{
		Interactor: usecase.PreImagesInteractor{
			PreImagesRepository: &database.PreImagesRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *PreImagesController) GetAll(c Context) {
	preimages, err := controller.Interactor.GetAllPreImages()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, preimages)
}
