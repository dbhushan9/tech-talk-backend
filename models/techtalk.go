package models

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type TechTalk struct {
	Id          string   `bson:"_id" json:"id"`
	Title       string   `bson:"title,omitempty" json:"title,omitempty"`
	Description string   `bson:"description,omitempty" json:"description,omitempty"`
	Speaker     string   `bson:"speaker,omitempty" json:"speaker,omitempty"`
	Tags        []string `bson:"tags,omitempty" json:"tags,omitempty"`
	Date        int64    `bson:"date,omitempty" json:"date,omitempty"`
	Archived    bool     `bson:"archived,omitempty" json:"archived,omitempty"`
}

func NewTechTalk(title string, description string, speaker string, tags []string, date int64, archived bool) *TechTalk {
	return &TechTalk{uuid.New().String(), title, description, speaker, tags, date, archived}
}

func (t *TechTalk) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}

func (t *TechTalk) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}
