package models

//Pet storage model for pet
type StaffGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Permissions string `json:"permissions"`
}

//PetService interface for Pet model
type StaffGroupService interface {
	CreateStaffGroup() error
	GetAllStaffGroup(StaffGroupID int) ([]*StaffGroup, error)
	GetStaffGroup(id int) (*StaffGroup, error)
	DeleteStaffGroup(id int) error
}
