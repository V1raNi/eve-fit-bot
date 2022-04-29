package models

import "gorm.io/gorm"

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
)

type Fit struct {
	gorm.Model
	Name        string
	Ship        string
	Link        string
	Activity    activityType
	Content     contentType
	Account     accountType
	Description string
}
