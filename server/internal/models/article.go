package models

import "time"


type Article struct {
    Content string `json:"content" validate:"required"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}
