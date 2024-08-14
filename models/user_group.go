package models

import "time"

type UserGroup struct {
    UserID    int       `json:"user_id" gorm:"primaryKey"`
    GroupID   int       `json:"group_id" gorm:"primaryKey"`
    AssignedAt time.Time `json:"assigned_at" gorm:"autoCreateTime"`
}