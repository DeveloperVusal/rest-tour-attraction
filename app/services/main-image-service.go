package services

import (
	"encoding/json"
	"net/http"

	"attrtour/app/http/dto"
	"attrtour/app/models"
	"attrtour/core"
)

type MainImageService struct {
	core.BaseService
	GetDto dto.GetMainImageDto
}

func (cs *MainImageService) GetById() {
	var mainImage models.MainImage

	cs.DBLink.Preload("Continent").Preload("Language").First(&mainImage, cs.GetDto.Id)

	jsonData, _ := json.Marshal(mainImage)

	cs.Wri.Write(jsonData)
}

func (mis *MainImageService) Upload(_dto dto.UploadMainImageDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		uploadService := &UploadImage{}
		uploadImagePath, _ := uploadService.Upload(mis.Wri, mis.Req, "/storage/uploads")

		mainImage := models.MainImage{
			LocationId:       _dto.LocationId,
			LanguageId:       _dto.LanguageId,
			UserId:           _dto.UserId,
			SourcePath:       uploadImagePath,
			ShortDescription: _dto.ShortDescription,
			IsMain:           &_dto.IsMain,
			IsVisible:        &_dto.IsVisible,
			IsArchive:        &_dto.IsArchive,
			IsDeleted:        &_dto.IsDeleted,
		}

		mis.DBLink.Create(&mainImage)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"main_image_id": mainImage.Id,
		})

		mis.Wri.WriteHeader(http.StatusCreated)
		mis.Wri.Write(jsonData)
	} else {
		mis.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (mis *MainImageService) Save(_dto dto.SaveMainImageDto) {
	validate := dto.DtoValidate{}

	if validate.Is(_dto) {
		var mainImage models.MainImage

		mis.DBLink.First(&mainImage, _dto.Id)

		if _dto.File != nil {
			uploadService := &UploadImage{}
			uploadImagePath, isUpload := uploadService.Upload(mis.Wri, mis.Req, "/storage/uploads")

			if isUpload {
				uploadService.Remove(mainImage.SourcePath)

				mainImage.SourcePath = uploadImagePath
			}
		}

		if len(_dto.ShortDescription) > 0 && _dto.ShortDescription != mainImage.ShortDescription {
			mainImage.ShortDescription = _dto.ShortDescription
		}

		if _dto.LocationId > 0 && _dto.LocationId != mainImage.LocationId {
			mainImage.LocationId = _dto.LocationId
		}

		if _dto.LanguageId > 0 && _dto.LanguageId != mainImage.LanguageId {
			mainImage.LanguageId = _dto.LanguageId
		}

		if _dto.UserId > 0 && _dto.UserId != mainImage.UserId {
			mainImage.UserId = _dto.UserId
		}

		if &_dto.IsMain != mainImage.IsMain {
			mainImage.IsMain = &_dto.IsMain
		}

		if &_dto.IsVisible != mainImage.IsVisible {
			mainImage.IsVisible = &_dto.IsVisible
		}

		if &_dto.IsArchive != mainImage.IsArchive {
			mainImage.IsArchive = &_dto.IsArchive
		}

		if &_dto.IsDeleted != mainImage.IsDeleted {
			mainImage.IsDeleted = &_dto.IsDeleted
		}

		mis.DBLink.Save(&mainImage)

		jsonData, _ := json.Marshal(map[string]interface{}{
			"status": "success",
		})

		mis.Wri.Write(jsonData)
	} else {
		mis.Wri.Write([]byte("Fields sent incorrectly\n"))
	}
}

func (mis *MainImageService) Remove() {
	var mainImage models.MainImage

	mis.DBLink.First(&mainImage, mis.GetDto.Id)
	mis.DBLink.Delete(&mainImage)

	uploadService := &UploadImage{}
	uploadService.Remove(mainImage.SourcePath)

	jsonData, _ := json.Marshal(map[string]interface{}{
		"status": "success",
	})

	mis.Wri.Write(jsonData)
}
