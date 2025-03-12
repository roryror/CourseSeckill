package domain

type Order struct {
	ID int `gorm:"primaryKey"`
	UserID int `gorm:"not null"`
	CourseID int `gorm:"not null"`
}

type Course struct {
	ID int `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Stock int `gorm:"not null"`
	MaxStock int `gorm:"not null"`
	MinStock int `gorm:"not null"`
}

type User struct {
	ID int `gorm:"primaryKey"`
	Name string
	Email string
	Password string
}
