package handlers

import (
	"io"
	"os"
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

func ImageFile(c *gin.Context) {
	Filename := c.Param("filename")

	if Filename == "" {
		http.Error(c.Writer, "Filename is not specified in url.", 400)
		return
	}

	fmt.Println("Client requests: " + Filename)

	fullPath := "./images/" + Filename

	Openfile, err := os.Open(fullPath)
	defer Openfile.Close()
	if err != nil {
		http.Error(c.Writer, "File not found.", 404)
		return
	}

	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)

	FileStat, _ := Openfile.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)

	c.Writer.Header().Set("Content-Disposition", "attachment; filename=" + Filename)
	c.Writer.Header().Set("Content-Type", FileContentType)
	c.Writer.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(c.Writer, Openfile)

	return
}
