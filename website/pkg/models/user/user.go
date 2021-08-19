package models

//User storage model for user
type User struct {
	ID        int    `json:"id"`
	GroupID   int    `json:"groupID"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Image     string `json:"image"`
	IP        string `json:"ip"`
	Status    bool   `json:"status"`
	DateAdded string `json:"dateAdded"`
}

//PetService interface for Pet model
type UserService interface {
	CreateUser() error
	GetUserByGroup(UserGroupID int) ([]*User, error)
	GetUser(id int) (*User, error)
	DeleteUser(id int) error
}
