package ecopedia

import (
	"time"
)

type Ecopedia struct {
	ID        int 
	Slug string
	Title     string
	SubTitle  string
	Description string
	SrcFile string
	Reference string
	FileName  []EcopediaImage `gorm:"foreignKey:EcopediaID"`
	CreatedAt time.Time
    UpdatedAt time.Time 
}

type EcopediaImage struct {
	ID        int      `gorm:"primaryKey"`
    FileName  string    
	EcopediaID int      
	CreatedAt time.Time 
    UpdatedAt time.Time 
}

