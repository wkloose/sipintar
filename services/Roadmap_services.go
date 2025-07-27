package services

import (
    "github.com/google/uuid"
    "github.com/wkloose/tempproject.git/initializers"
    "github.com/wkloose/tempproject.git/models"
)

type RoadmapItem struct {
    ID          uuid.UUID `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Unlocked    bool      `json:"unlocked"`
    Percent     float64   `json:"percent"`
    Status      string    `json:"status"` // "locked", "in_progress", "completed"
}

func GetRoadmapForUser(userID uuid.UUID) ([]RoadmapItem, error) {
    var materials []models.Material
    err := initializers.DB.Find(&materials).Error
    if err != nil {
        return nil, err
    }

    var progressList []models.LearningProgress
    _ = initializers.DB.Where("user_id = ?", userID).Find(&progressList)

    progressMap := make(map[uuid.UUID]models.LearningProgress)
    for _, p := range progressList {
        progressMap[p.MaterialID] = p
    }

    roadmap := []RoadmapItem{}

    for i, m := range materials {
        p, found := progressMap[m.ID]

        item := RoadmapItem{
            ID:          m.ID,
            Title:       m.Title,
            Description: m.Description,
            Percent:     0,
            Unlocked:    i == 0 || found, 
            Status:      "locked",
        }

        if found {
            item.Percent = p.Percent
            if p.Percent >= 90 {
                item.Status = "completed"
            } else {
                item.Status = "in_progress"
            }
        } else if i == 0 {
            item.Status = "unlocked"
        }

        roadmap = append(roadmap, item)
    }

    return roadmap, nil
}
func GetTotalProgress(userID uuid.UUID) (float64, error) {
    var materials []models.Material
    err := initializers.DB.Find(&materials).Error
    if err != nil {
        return 0, err
    }

    var progressList []models.LearningProgress
    _ = initializers.DB.Where("user_id = ?", userID).Find(&progressList)

    progressMap := make(map[uuid.UUID]float64)
    for _, p := range progressList {
        progressMap[p.MaterialID] = p.Percent
    }

    total := 0.0
    for _, m := range materials {
        total += progressMap[m.ID] 
    }

    if len(materials) == 0 {
        return 0.0, nil
    }

    return total / float64(len(materials)), nil
}

