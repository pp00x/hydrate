package handler

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
	"github.com/pp00x/hydrate/internal/model"
    "github.com/pp00x/hydrate/internal/repository"
    "github.com/pp00x/hydrate/internal/service"
    "github.com/pp00x/hydrate/pkg/util"
)

var waterIntakeService service.WaterIntakeService

func init() {
    db := util.GetDB()
    intakeRepo := repository.NewWaterIntakeRepository(db)
    waterIntakeService = service.NewWaterIntakeService(intakeRepo)
}

type WaterIntakeInput struct {
    Amount  float64 `json:"amount" binding:"required,gt=0"`
    TakenAt string  `json:"taken_at" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

func CreateWaterIntake(c *gin.Context) {
    var input WaterIntakeInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    takenAt, err := time.Parse(time.RFC3339, input.TakenAt)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
        return
    }

    intake := model.WaterIntake{
        UserID:  userID.(uint),
        Amount:  input.Amount,
        TakenAt: takenAt,
    }

    if err := waterIntakeService.Create(&intake); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Water intake recorded"})
}

func GetWaterIntakes(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    startDateStr := c.Query("start_date")
    endDateStr := c.Query("end_date")

    var startDate, endDate time.Time
    var err error

    if startDateStr == "" {
        startDate = time.Now().AddDate(0, 0, -7) // Default to last 7 days
    } else {
        startDate, err = time.Parse("2006-01-02", startDateStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format"})
            return
        }
    }

    if endDateStr == "" {
        endDate = time.Now()
    } else {
        endDate, err = time.Parse("2006-01-02", endDateStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format"})
            return
        }
    }

    intakes, err := waterIntakeService.GetByUserID(userID.(uint), startDate, endDate)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": intakes})
}