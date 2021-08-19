package service

import (
	db "github.com/MDPaun/go-start/tree/main/website/internal/staff/repo/postgres"

	models "github.com/MDPaun/go-start/tree/main/website/pkg/models/staff"
)

var us db.StaffService
var staff db.Staff

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

//GetStaff get one User
func GetStaff(id int) (*models.Staff, error) {
	us = &staff
	staff, err := us.GetStaff(id)
	if err != nil {
		return nil, err
	}
	return staff, nil
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
