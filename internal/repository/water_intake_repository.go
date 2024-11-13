package repository

import (
    "time"

    "gorm.io/gorm"
    "github.com/pp00x/hydrate/internal/model"
)

type WaterIntakeRepository interface {
    Create(intake *model.WaterIntake) error
    GetByUserID(userID uint, startDate, endDate time.Time) ([]model.WaterIntake, error)
}

type waterIntakeRepository struct {
    db *gorm.DB
}

func NewWaterIntakeRepository(db *gorm.DB) WaterIntakeRepository {
    return &waterIntakeRepository{db}
}

func (r *waterIntakeRepository) Create(intake *model.WaterIntake) error {
    return r.db.Create(intake).Error
}

func (r *waterIntakeRepository) GetByUserID(userID uint, startDate, endDate time.Time) ([]model.WaterIntake, error) {
    var intakes []model.WaterIntake
    result := r.db.Where("user_id = ? AND taken_at BETWEEN ? AND ?", userID, startDate, endDate).Find(&intakes)
    return intakes, result.Error
}