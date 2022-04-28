package data

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model `json:"-"`
	ID         uint   `gorm:"primary_key" json:"id"`
	Text       string `json:"text" validate:"required,gte=1"`
	Day        string `json:"day"`
	Reminder   bool   `json:"reminder"`
}

func (t *Task) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(t)
}

func (t *Task) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(t)
}

type Tasks []*Task

func (tl *Tasks) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(tl)
}

func (t *Task) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
