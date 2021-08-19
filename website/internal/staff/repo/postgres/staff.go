package postgres

import (
	"database/sql"
	"errors"

	client "github.com/MDPaun/go-start/tree/main/website/internal/database/postgres"

	models "github.com/MDPaun/go-start/website/pkg/models/user"
)

//User use from models
type User models.User

//UsertService use from models
type UserService models.UserService

//GetUser get one User
func (u User) GetUser(id int) (*models.User, error) {
	row := client.DbClient.QueryRow("select user.id, user.user_group.id, user.email, user.password, user.firstname, user.lastname, user.image, user.ip, user.status, user.date_added, user_group.name from user inner join user_group on user.user_group.id = user_group.id where pet.id=$1;", id)
	switch err := row.Scan(&u.ID, &u.GroupID, &u.Email, &u.Password, &u.FirstName, &u.LastName, &u.Image, &u.IP, &u.Status, &u.Date_added, &u.User_group.name); err {
	case sql.ErrNoRows:
		return nil, errors.New("No matching records")
	case nil:
		category := models.User(u)
		return &category, nil
	default:
		return nil, err
	}

}

// //GetPetByCategory based on category, get all Pets
// func (p Pet) GetPetByCategory(categoryID int) ([]*models.Pet, error) {
// 	var pets []*models.Pet
// 	rows, err := client.DbClient.Query("select pet.id,pet.pet_name,pet.age,pet.description,pet.category_id,category.category_name,pet.breed_id,breed.breed_name,pet.location_id,location.location_name from pet inner join category on pet.category_id = category.id inner join breed on pet.breed_id=breed.id inner join location on pet.location_id = location.id where category.id=$1;", categoryID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		err = rows.Scan(&p.ID, &p.Name, &p.Age, &p.Description, &p.PetCategory.ID, &p.PetCategory.CategoryName, &p.PetBreed.ID, &p.PetBreed.BreedName, &p.PetLocation.ID, &p.PetLocation.LocationName)
// 		pm := models.Pet(p)
// 		pets = append(pets, &pm)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	if rows.Err() != nil {
// 		return nil, err
// 	}

// 	return pets, nil
// }

// //CreatePet create a Pet
// func (p *Pet) CreatePet() error {
// 	stmt, err := client.DbClient.Prepare("INSERT INTO pet (name,age,image_url,description,breed_id,category_id,location_id) VALUES($1,$2,$3,$4,$5,$6,$7);")
// 	if err != nil {
// 		return err
// 	}
// 	//closing the statement to prevent memory leaks
// 	defer stmt.Close()
// 	_, err = stmt.Exec(p.Name, p.Age, p.ImageURL, p.Description, p.PetBreed.ID, p.PetCategory.ID, p.PetLocation.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //DeletePet delete a Pet
// func (p *Pet) DeletePet(id int) error {
// 	stmt, err := client.DbClient.Prepare("DELETE from pet where id=$1")
// 	if err != nil {
// 		return err
// 	}
// 	//closing the statement to prevent memory leaks
// 	defer stmt.Close()
// 	_, err = stmt.Exec(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
