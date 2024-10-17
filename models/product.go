package models

import (
	"errors"
	"html"
	"mime/multipart"
	"strings"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name 		string		`json:"name" validate:"required"`
	Description string		`json:"description" validate:"required"`
	Price 		float64		`json:"price" validate:"required"`
	Quantity 	int			`json:"qty" validate:"required"`
	Image		string		`validate:"required"`
	CategoryID 	uint   
    Category   	Category 	`gorm:"foreignKey:CategoryID"`
}

func (f *Product) ValidateFile(file *multipart.FileHeader) error {
	const maxFileSize = 10 << 20
	allowedTypes := []string{"image/jpg","image/png"}
	
	if file.Size > maxFileSize{
		return errors.New("file size too big max is 10MB")
	}
	fileType := file.Header.Get("Content-Type")
	if !isFileTypeAllowed(fileType,allowedTypes) {
		return errors.New("unsupported file")
	}
	return nil
}

func isFileTypeAllowed(fileType string, allowedTypes []string) bool {
    for _, allowed := range allowedTypes {
        if strings.EqualFold(fileType, allowed) {
            return true
        }
    }
    return false
}

func (p *Product) BeforeSave(tx *gorm.DB) (err error){
	p.Name = strings.TrimSpace(p.Name)
	p.Description = strings.TrimSpace(p.Description)
	
	p.Name = html.EscapeString(p.Name)
	p.Description = html.EscapeString(p.Description)

	return nil
}

