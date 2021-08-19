package models

//Pet storage model for pet
type UserGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Permissions string `json:"permissions"`
}

//PetService interface for Pet model
type UserGroupService interface {
	CreateUserGroup() error
	GetAllUserGroup(UserGroupID int) ([]*UserGroup, error)
	GetUserGroup(id int) (*UserGroup, error)
	DeleteUserGroup(id int) error
}
