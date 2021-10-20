package models

//Staff storage model for Staff
type Staff struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FullName  string `json:"fullname"`
	Image     string `json:"image"`
	Status    bool   `json:"status"`
	DateAdded string `json:"dateAdded"`
}

//PetService interface for Pet model
type StaffRepository interface {
	GetStaff() error
	// FindByID(id int) (*Staff, error)
	// DeleteStaff(id int) error
	// CreateStaff() error
	// GetStaffByGroup(StaffGroupID int) ([]*Staff, error)
}
