package main

import (
	"bytes"
	"image"
	"image/png"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
)

func generateQRCode(c *gin.Context) {
	sizeStr := c.Query("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size"})
		return
	}

	data := c.Query("data")

	qr, err := qrcode.Encode(data, qrcode.Medium, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	qrImage := bytes.NewBuffer(qr)
	img, _, err := image.Decode(qrImage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resized := resize.Resize(100, 0, img, resize.Lanczos3)

	var buf bytes.Buffer
	if err := png.Encode(&buf, resized); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "image/png", buf.Bytes())
}

func main() {
	router := gin.Default()

	router.GET("/qr", generateQRCode)

	router.Run(":8080")
}
