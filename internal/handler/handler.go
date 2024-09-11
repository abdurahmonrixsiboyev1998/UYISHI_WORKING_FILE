package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)
const uploadPath = "./Downloads"

const maxUploadSize = 100 * 1024 * 1024

var allowedFileTypes = map[string]bool{
	".txt": true,
	".jpg": true,
	".png": true,
	".pdf": true,
	".md": 	true,
}

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Faqat POST so'rovlari qabul qilinadi", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		http.Error(w, "Fayl hajmi juda katta!", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Faylni yuklashda xato", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileExt := strings.ToLower(filepath.Ext(handler.Filename))
	if !allowedFileTypes[fileExt] {
		http.Error(w, "Bu turdagi fayllarga ruxsat berilmagan", http.StatusBadRequest)
		return
	}

	dst, err := os.Create(filepath.Join(uploadPath, handler.Filename))
	if err != nil {
		http.Error(w, "Faylni saqlashda xato", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Faylni nusxalashda xato", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Fayl muvaffaqiyatli yuklandi: %s\n", handler.Filename)))
}



func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {

	fileName := strings.TrimPrefix(r.URL.Path, "./Downloads/")
	filePath := filepath.Join(uploadPath, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "Fayl topilmadi", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filePath)
}
