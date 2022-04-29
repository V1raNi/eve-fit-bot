package models

import (
	"gorm.io/gorm"
	"time"
)

type activityType string

const (
	PvP   activityType = "PvP"
	PvE   activityType = "PvE"
	Abyss activityType = "Abyss"
)

type accountType string

const (
	Alpha accountType = "Alpha"
	Omega accountType = "Omega"
)

type contentType string

const (
	WH         contentType = "Wormholes"
	Anomalies  contentType = "Anomalies"
	Solo       contentType = "Solo"
	SmallScale contentType = "Small-scale"
	Event      contentType = "Event"
	Scan       contentType = "Scan"
	ESS        contentType = "ESS"
)

type Fit struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name        string         `json:"name"`
	Ship        string         `json:"ship"`
	Link        string         `json:"link"`
	Activity    activityType   `json:"activity"`
	Content     contentType    `json:"content"`
	Account     accountType    `json:"account"`
	Description string         `json:"description"`
}
