package helper

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
)

func UploadImageMovies(c *gin.Context, fileHeader *multipart.FileHeader) (string, error) {
	nameImage := GenerateUUID() + filepath.Ext(fileHeader.Filename)

	err := c.SaveUploadedFile(fileHeader, StorageImage+"movies/"+nameImage)
	if err != nil {
		return nameImage, err
	}

	return nameImage, nil
}

func GetImageMovies(c *gin.Context, newImage string) string {
	return c.Request.Host + "/" + StorageImage + "movies/" + newImage
}
