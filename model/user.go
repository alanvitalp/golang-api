package model

type User struct {
	ID					string   `json:"id" gorm:"primary_key"`
	Username		string	`json:"username" gorm:"not null;unique"`
	FirstName   string  `json:"firstName" gorm:"not null"`
	LastName    string  `json:"lastName" gorm:"not null"`
	Email       string  `json:"email" gorm:"not null"`
	Password    string  `json:"password" gorm:"not null"`
	Phone			 	string  `json:"phone" gorm:"not null"`
	UserStatus	int  		`json:"userStatus" gorm:"not null"`
}