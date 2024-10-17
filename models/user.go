package models

import (
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string    
	Email     string  `gorm:"unique"`  
	Password  string    
	Roles 	  []Role  `gorm:"many2many:user_roles;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:role_id"`
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time
}

// type UserRole struct{
// 	UserID  int `gorm:"foreignKey:id;References:id;index"`
//   	RoleID int `gorm:"foreignKey:id;References:id;index"`
//   	CreatedAt time.Time
//   	DeletedAt gorm.DeletedAt
// }

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// trim
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)

	// escape (&,>,<,"")
	u.Name = html.EscapeString(u.Name)
	u.Email = html.EscapeString(u.Email)
	u.Password = html.EscapeString(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//var role Role
	u.Roles = []Role{
		{ID:2 },
	}
	u.Password = string(hashedPassword)
	return nil
}

// func (ur *UserRole) BeforeSave(tx *gorm.DB) error {
// 	var u *User
// 	user := &UserRole{
// 		UserID: int(u.Id),
// 		RoleID: 2,
// 	}
// 	if err := tx.Create(user).Error; err != nil{
// 		return err
// 	}
// 	return nil
// }