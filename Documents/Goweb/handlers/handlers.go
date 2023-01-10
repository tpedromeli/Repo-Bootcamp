package handlers

import (
	"Goweb/models"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var products, _ = getProductsStruct()

func openJsonFile() (jsonFile *os.File, err error) {

	jsonFile, err = os.Open("products.json")

	if err != nil {
		return
	}
	return
}

func getProductsStruct() (products models.Products, err error) {

	jsonFile, err := openJsonFile()

	if err != nil {
		return
	}
	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return
	}

	if err = json.Unmarshal(byteValue, &products.Products); err != nil {
		return
	}

	defer jsonFile.Close()

	return
}

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func List(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, products)
}
