package entity

import (
	"encoding/json"

	"newser.app/internal/domain/value"
	"newser.app/shared"
)

type Image struct {
	ID    ID         `json:"id"`
	URL   value.Link `json:"url"`
	Title string     `json:"title"`
}

func NewImage(url, title string) *Image {
	validLink, err := value.NewLink(url)
	if err != nil {
		valErr, ok := err.(shared.AppError)
		if ok {
			valErr.Print()
		}
		return nil
	}
	return &Image{
		ID:    NewID(),
		URL:   validLink,
		Title: title,
	}
}

func (i Image) JSON() []byte {
	json, _ := json.MarshalIndent(i, "", "  ")
	return json
}

func (i Image) String() string {
	return string(i.JSON())
}
