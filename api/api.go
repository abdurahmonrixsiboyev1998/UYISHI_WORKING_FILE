package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"working-file/internal/handler"
)


var uploadPath = "./Downloads"
func ConnApi() {
	http.HandleFunc("/upload", handler.UploadFileHandler)
	http.HandleFunc("/download/", handler.DownloadFileHandler)

	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server Running on  http://localhost:8888")
	http.ListenAndServe(":8888", nil)
}
