package models

import "time"

type Group struct {
    ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
    Name        string    `json:"name" gorm:"unique;not null"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}