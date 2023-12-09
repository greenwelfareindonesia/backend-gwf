package ecopedia

import (
	"greenwelfare/user"
	"time"
)

type Ecopedia struct {
	ID        int 
	Judul     string
	Subjudul  string
	Deskripsi string
	Srcgambar string
	Referensi string
	// Like []IsLike
	FileName       []EcopediaImage `gorm:"foreignKey:EcopediaID"`
	Comment []Comment `gorm:"foreignKey:EcopediaId"`
	// IsLike []IsLike `gorm:"foreignKey:EcopediaId"`
	CreatedAt time.Time
    UpdatedAt time.Time 
}

type EcopediaImage struct {
	ID        int      `gorm:"primaryKey"`
    FileName  string    
	EcopediaID  int      
	CreatedAt time.Time 
    UpdatedAt time.Time 
}

type Comment struct {
	ID int
	UserId int
	EcopediaId int
	Comment string
	User user.User
	CreatedAt time.Time
    UpdatedAt time.Time
}

// type IsLike struct {
// 	ID int
// 	UserId int
// 	EcopediaId int
// 	IsLike bool
// 	User user.User
// 	CreatedAt time.Time
//     UpdatedAt time.Time
// }