package service

import (
	db "github.com/MDPaun/go-start/tree/main/website/internal/database/postgres"

	models "github.com/MDPaun/go-start/website/pkg/models/staff"
)

var us db.StaffService
var staf db.Staff

// //CreatePet create a Pet
// func CreatePet(p *db.Pet) error {
// 	ps = p
// 	err = ps.CreatePet()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //DeletePet deletes a Pet
// func DeletePet(id int) error {
// 	ps = &pet
// 	err = ps.DeletePet(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//GetUser get one User
func GetStaff(id int) (*models.User, error) {
	us = &user
	user, err := us.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// //GetPetByCategory get Pet based on category
// func GetPetByCategory(categoryID int) ([]*models.Pet, error) {
// 	ps = &pet
// 	pets, err := ps.GetPetByCategory(categoryID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pets, nil
// }
