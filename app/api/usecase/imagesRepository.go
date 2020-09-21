package usecase

import (
	"../domain"
)

type (
	ImagesRepository interface {
		// 	Add(domain.Upload) (domain.Upload, error)
		// 	Get(int) (domain.Upload, error)
		// 	Delete(domain.Upload) (domain.Upload, error)
		GetAll() (domain.Uploads, error)
	}
	PreImagesRepository interface {
		//Add(domain.Preupload) (domain.Preupload, error)
		// Get(int) (domain.Preupload, error)
		// Delete(domain.Preupload) (domain.Preupload, error)
		GetAll() (domain.Preuploads, error)
	}
)
