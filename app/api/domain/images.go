package domain

import (
	"time"
)

type (

	// Preupload model
	Preupload struct {
		ID          int
		Title       string
		Path        string
		ShotDate    *time.Time
		CreatedDate *time.Time
		DeletedDate *time.Time
	}

	//Upload model
	Upload struct {
		ID          int
		Title       string
		Path        string
		ShotDate    *time.Time
		CreatedDate *time.Time
		DeletedDate *time.Time
	}

	Preuploads []Preupload

	Uploads []Upload

// // PreImageController struct
// PreImageController struct{}
// // ImageController struct
// ImageController struct{}
)
