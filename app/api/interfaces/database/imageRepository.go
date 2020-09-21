package database

import (
	"../../domain"
	"time"
)

// Upload

type ImagesRepository struct {
	SqlHandler
}

func (repo *ImagesRepository) GetAll() (Uploads domain.Uploads, err error) {
	rows, err := repo.Query("SELECT * FROM upload")
	defer rows.Close()
	if err != nil {
		return
	}
	var result domain.Uploads
	for rows.Next() {
		var id int
		var title string
		var path string
		var shotDate *time.Time
		var createdDate *time.Time
		var deletedDate *time.Time

		if err := rows.Scan(
			&id,
			&title,
			&path,
			&shotDate,
			&createdDate,
			&deletedDate,
		); err != nil {
			continue
		}
		upload := domain.Upload{
			ID:          id,
			Title:       title,
			Path:        path,
			ShotDate:    shotDate,
			CreatedDate: createdDate,
			DeletedDate: deletedDate,
		}
		result = append(result, upload)
	}

	return result, err
}

// PreUpload

type PreImagesRepository struct {
	SqlHandler
}

func (repo *PreImagesRepository) GetAll() (preuploads domain.Preuploads, err error) {
	rows, err := repo.Query("SELECT * FROM preupload")
	defer rows.Close()
	if err != nil {
		return
	}
	var result domain.Preuploads
	for rows.Next() {
		var id int
		var title string
		var path string
		var shotDate *time.Time
		var createdDate *time.Time
		var deletedDate *time.Time

		if err := rows.Scan(
			&id,
			&title,
			&path,
			&shotDate,
			&createdDate,
			&deletedDate,
		); err != nil {
			continue
		}
		preupload := domain.Preupload{
			ID:          id,
			Title:       title,
			Path:        path,
			ShotDate:    shotDate,
			CreatedDate: createdDate,
			DeletedDate: deletedDate,
		}
		result = append(result, preupload)
	}

	return result, err
}
