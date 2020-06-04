package history

import "time"

type History struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     []byte    `json:"content"`
	Date        time.Time `json:"date"`
}

type History_image struct {
	ImageId       string `json:"image_id"`
	History_image string `json:"history_image"`
	Description   string `json:"description"`
}
