package service

import (
    "time"

    "github.com/pp00x/hydrate/internal/model"
    "github.com/pp00x/hydrate/internal/repository"
)

type WaterIntakeService interface {
    Create(intake *model.WaterIntake) error
    GetByUserID(userID uint, startDate, endDate time.Time) ([]model.WaterIntake, error)
}

type waterIntakeService struct {
    intakeRepo repository.WaterIntakeRepository
}

func NewWaterIntakeService(intakeRepo repository.WaterIntakeRepository) WaterIntakeService {
    return &waterIntakeService{intakeRepo}
}

func (s *waterIntakeService) Create(intake *model.WaterIntake) error {
    return s.intakeRepo.Create(intake)
}

func (s *waterIntakeService) GetByUserID(userID uint, startDate, endDate time.Time) ([]model.WaterIntake, error) {
    return s.intakeRepo.GetByUserID(userID, startDate, endDate)
}