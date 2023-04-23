package models

import "time"

type ImageMetadata struct {
	ImageName   string
	ImageUrl    string
	ContentType string
	Size        int64
	CreatedAt   time.Time
}
