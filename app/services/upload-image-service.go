package services

import (
	"crypto/rand"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type UploadImage struct{}

func (ui *UploadImage) Upload(w http.ResponseWriter, r *http.Request, uploadPath string) (string, bool) {
	file, _, err := r.FormFile("file")

	if err != nil {
		ui.RenderError(w, "INVALID_FILE", http.StatusBadRequest)

		return "", false
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)

	if err != nil {
		ui.RenderError(w, "INVALID_FILE", http.StatusBadRequest)

		return "", false
	}

	detectedFileType := http.DetectContentType(fileBytes)
	switch detectedFileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		ui.RenderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
		return "", false
	}

	fileName := ui.RandToken(12)
	fileEndings, err := mime.ExtensionsByType(detectedFileType)

	if err != nil {
		ui.RenderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)

		return "", false
	}
	projectPath := "D:/Development/golang/rest-tour-attraction"
	sourcePath := strings.ReplaceAll(filepath.Join(uploadPath, fileName+fileEndings[0]), "\\", "/")
	savePath := strings.ReplaceAll(filepath.Join(projectPath, uploadPath, fileName+fileEndings[0]), "\\", "/")
	// fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

	// write file
	newFile, err := os.Create(savePath)

	if err != nil {
		fmt.Println(err)
		ui.RenderError(w, "CANT_WRITE_FILE_1", http.StatusInternalServerError)

		return "", false
	}

	defer newFile.Close() // idempotent, okay to call twice

	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		ui.RenderError(w, "CANT_WRITE_FILE_2", http.StatusInternalServerError)
		return "", false
	}

	return sourcePath, true
}

func (ui *UploadImage) Remove(uploadPath string) {
	projectPath := "D:/Development/golang/rest-tour-attraction"
	buildPath := strings.ReplaceAll(filepath.Join(projectPath, uploadPath), "\\", "/")

	os.Remove(buildPath)
}

func (ui *UploadImage) RenderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func (ui *UploadImage) RandToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
