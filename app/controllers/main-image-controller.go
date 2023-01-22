package controllers

import (
	"github.com/gorilla/schema"

	"attrtour/app/http/dto"
	"attrtour/app/services"
	"attrtour/core"
)

type MainImageController core.BaseController

func (mic *MainImageController) GetById(id uint) {
	mis := services.MainImageService{
		GetDto: dto.GetMainImageDto{
			Id: id,
		},
	}
	mis.RequestInit(mic.DBLink, mic.Wri, mic.Req)

	mis.GetById()
}

func (mic *MainImageController) Upload() {
	var decoder = schema.NewDecoder()
	var dtoData dto.UploadMainImageDto

	mic.Req.ParseMultipartForm(128 << 20)
	decoder.Decode(&dtoData, mic.Req.Form)

	_, fileHeader, _ := mic.Req.FormFile("file")
	dtoData.File = fileHeader

	mis := services.MainImageService{}
	mis.RequestInit(mic.DBLink, mic.Wri, mic.Req)

	mis.Upload(dtoData)
}

func (mic *MainImageController) Save() {
	var decoder = schema.NewDecoder()
	var dtoData dto.SaveMainImageDto

	mic.Req.ParseMultipartForm(128 << 20)
	decoder.Decode(&dtoData, mic.Req.Form)

	_, fileHeader, _ := mic.Req.FormFile("file")
	dtoData.File = fileHeader

	mis := services.MainImageService{}
	mis.RequestInit(mic.DBLink, mic.Wri, mic.Req)

	mis.Save(dtoData)
}

func (mic *MainImageController) Remove(id uint) {
	mis := services.MainImageService{
		GetDto: dto.GetMainImageDto{
			Id: id,
		},
	}
	mis.RequestInit(mic.DBLink, mic.Wri, mic.Req)

	mis.Remove()
}
