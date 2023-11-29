package models

type Tag struct {
	Name string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
	Id string `json:"id,omitempty"`

	Count int `json:"count,omitempty"`
}
