package models

import (
	"crypto/rand"
	"math/big"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        			uint    `gorm:"primaryKey"`
	Name      			string    	`gorm:"type:varchar(255);not null"`
	Email     			string  	`gorm:"unique"`  
	Password  			string    	`gorm:"not null"`
	Pin		  			string
	Profile   			string	
	Token				[]Token
	CreatedAt 			time.Time // Automatically managed by GORM for creation time
	UpdatedAt 			time.Time
}

type Token struct {
	gorm.Model
	TokenString		string		`gorm:"unique:not null"`
	ExpiredAt		time.Time	`gorm:"not null"`
	IsUsed			bool		`gorm:"default:false"`
	UserID			uint
	User			User
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return err
		}
		b[i] = characters[num.Int64()]
	}
	u.Pin = string(b)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	u.Password = string(hashedPassword)
	return nil
}

func (t *Token) BeforeSave(tx *gorm.DB) (err error) {
	
	expiryTime := time.Now().Add(24 * time.Hour)
	//t.TokenString = genToken
	t.ExpiredAt = expiryTime
	return nil
}