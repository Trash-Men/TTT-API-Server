package models

import "time"

type TrashCan struct {
	tableName  struct{}  `pg:"trash_can"`
	Photo_url  string    `json:"photo_url"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Area       string    `json:"area"`
	Created_at time.Time `json:"created_at"`
}
