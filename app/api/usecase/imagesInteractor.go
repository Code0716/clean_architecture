package usecase

import "../domain"

type ImagesInteractor struct {
	ImagesRepository ImagesRepository
}

type PreImagesInteractor struct {
	PreImagesRepository PreImagesRepository
	//ImagesRepository    ImagesRepository
}

// // CreateFileInfo save  upload info
// func (interactor *ImagesInteractor) CreateFileInfo(f domain.Upload) (file domain.Upload, err error) {
// 	file, err = interactor.ImagesRepository.Add(f)
// 	return
// }

// // GetFileInfo 個々のファイルを取得
// func (interactor *ImagesInteractor) GetFileInfo(id int) (file domain.Upload, err error) {
// 	file, err = interactor.ImagesRepository.Get(id)
// 	return
// }

// // DeleteFileInfo 個々のファイルを削除
// func (interactor *ImagesInteractor) DeleteFileInfo(f domain.Upload) (file domain.Upload, err error) {
// 	file, err = interactor.ImagesRepository.Delete(f)
// 	return
// }

// GetAllFileInfo uplad全件取得
func (interactor *ImagesInteractor) GetAllImages() (files domain.Uploads, err error) {
	files, err = interactor.ImagesRepository.GetAll()
	return
}

// CreatePreFileInfo save  preupload infogit
// func (interactor *ImagesInteractor) CreatePreFileInfo(f domain.Preupload) (file domain.Preupload, err error) {
// 	file, err = interactor.PreImagesRepository.Add(f)
// 	return
// }

// // GetPreFileInfo 個々のファイルを取得
// func (interactor *ImagesInteractor) GetPreFileInfo(id int) (file domain.Preupload, err error) {
// 	file, err = interactor.PreImagesRepository.Get(id)
// 	return
// }

// // DeletePreFileInfo 個々のファイルを削除
// func (interactor *ImagesInteractor) DeletePreFileInfo(f domain.Preupload) (file domain.Preupload, err error) {
// 	file, err = interactor.PreImagesRepository.Delete(f)
// 	return
// }

// GetAllPreImages preuplad全件取得
func (interactor *PreImagesInteractor) GetAllPreImages() (files domain.Preuploads, err error) {
	files, err = interactor.PreImagesRepository.GetAll()
	return
}
