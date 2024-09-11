package models

import "time"

type FileMetaDate struct {
	Size         int32
	FileName     string
	UploadedTime time.Time
}
