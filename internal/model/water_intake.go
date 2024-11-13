package model

import "time"

type WaterIntake struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `gorm:"not null" json:"user_id"`
    Amount    float64   `gorm:"not null" json:"amount"` // in milliliters
    TakenAt   time.Time `gorm:"not null" json:"taken_at"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}