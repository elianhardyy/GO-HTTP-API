package repositories

// import (
// 	"server/models"

// 	"gorm.io/gorm"
// )

// type GroupRepository interface {
// 	CreateGroup(group models.Group)(models.Group,error)
// 	AddMember()
// 	AddDescription()
// 	DeleteGroup()
// }

// type groupRepository struct{
// 	DB *gorm.DB
// }

// func NewGroupRepository(DB *gorm.DB) GroupRepository {
// 	return &groupRepository{
// 		DB: DB,
// 	}
// }

// func (g *groupRepository)CreateGroup(group models.Group)(models.Group,error){
// 	if err := g.DB.Create(&group).Error; err != nil {

// 	}
// }