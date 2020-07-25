package model

import (
	"github.com/jinzhu/gorm"
)

// Incident ...
type Incident struct {
	gorm.Model
	Name        string ` json:"name"`
	Description string ` json:"description"`
	Priority    int    ` gorm:"default:'3'" json:"priority"`
	Type        string ` gorm:"default:'query'" json:"type"`
	Acknowlege  string ` gorm:"default:'false'" json:"acknowlege"`
	Status      string ` gorm:"default:'open'" json:"status"`
	Severity    string ` gorm:"default:'medium'" json:"severity"`
	ReportedBy  string ` json:"reportedBy"`
}

// Find ...
type Find struct {
	ID          int    ` json:"id"`
	Name        string ` json:"name"`
	Description string ` json:"description"`
	Priority    int    ` json:"priority"`
	Type        string ` json:"type"`
	Acknowlege  string ` json:"acknowlege"`
	Status      string ` json:"status"`
	Severity    string ` json:"severity"`
	ReportedBy  string ` json:"reportedBy"`
}

// GetLastIncident ...
func GetLastIncident() *Incident {
	incident := &Incident{}
	GetDB().Last(&incident)
	return incident
}

// InsertIncident ...
func InsertIncident(incident *Incident) {
	GetDB().Create(incident)
}

// UpdateIncidentByID ...
func UpdateIncidentByID(incident *Incident, id int64) {
	GetDB().Save(&incident)

}

// GetIncidentByID ...
func GetIncidentByID(id int64) *Incident {
	incident := &Incident{}
	GetDB().Where([]int64{id}).First(&incident)
	return incident

}
