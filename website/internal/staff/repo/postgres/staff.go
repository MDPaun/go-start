package postgres

import (
	"fmt"

	client "github.com/MDPaun/go-start/tree/main/website/internal/database/postgres"
	models "github.com/MDPaun/go-start/tree/main/website/pkg/models/staff"
)

//Staff use from models
type Staff models.Staff

//StafftService use from models
type StaffService models.StaffService

//GetStaff get one User
func (s Staff) GetStaff(id int) (*models.Staff, error) {
	db := client.ConnectDB()
	defer db.Close()
	fmt.Println("Service")
	// row := client.DbClient.QueryRow("SELECT staff.id, staff.staff_group_id, staff.email, staff.password, staff.firstname, staff.lastname, staff.image, staff.ip, staff.status, staff.date_added FROM staff where staff.id=$1;", 1) //, staff_group.name from user inner join name on staff.staff_group_id, = staff_group.id
	row := db.QueryRow("SELECT * FROM staff WHERE staff_group_id = 1;", id) //, staff_group.name from user inner join name on staff.staff_group_id, = staff_group.id

	err := row.Scan(&s.ID, &s.GroupID, &s.Email, &s.Password, &s.Salt, &s.FirstName, &s.LastName, &s.Image, &s.IP, &s.Status, &s.DateAdded)
	if err != nil {
		fmt.Println(err)
	}
	return nil, err
	// switch err := row.Scan(&s.ID, &s.GroupID, &s.Email, &s.Password, &s.FirstName, &s.LastName, &s.Image, &s.IP, &s.Status, &s.DateAdded); err {
	// case sql.ErrNoRows:
	// 	return nil, errors.New("no matching records")
	// case nil:
	// 	category := models.Staff(s)
	// 	return &category, nil
	// default:
	// 	return nil, err
	// }

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

//CreatePet create a Pet
// func (p *Pet) CreateStaff() error {
// 	// stmt, err := client.DbClient.Prepare("INSERT INTO pet (name,age,image_url,description,breed_id,category_id,location_id) VALUES($1,$2,$3,$4,$5,$6,$7);")
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
