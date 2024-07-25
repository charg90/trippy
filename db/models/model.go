package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"deletedAt"`
	FirstName string     `json:"firstName"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	Plans     []Plan     `gorm:"foreignKey:UserId" json:"plans"`
}

type Plan struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"deletedAt"`
	Name      string     `json:"name"`
	StartDate time.Time  `json:"startDate"`
	EndDate   time.Time  `json:"endDate"`
	UserId    uuid.UUID  `json:"userId"`
	Trips     []Trip     `gorm:"foreignKey:PlanId" json:"trips"`
}

type Trip struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"deletedAt"`
	Name      string     `json:"name"`
	StartDate time.Time  `json:"startDate"`
	EndDate   time.Time  `json:"endDate"`
	PlanId    uuid.UUID  `json:"planId"`
	Location  string     `json:"location"`
	Activities []Activity `gorm:"foreignKey:TripId" json:"activities"`
}

type Activity struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"deletedAt"`
	Name      string     `json:"name"`
	StartDate time.Time  `json:"startDate"`
	ActivityType  ActivityType `gorm:"foreignKey:ActivityTypeID" json:"activityType"`
	ActivityTypeID uuid.UUID    `json:"activityTypeId"`
	Price     float64    `json:"price"`
	TripId    uuid.UUID  `json:"tripId"`
}

type ActivityType struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Type string    `json:"type"`
	Image string    `json:"image"` 
}