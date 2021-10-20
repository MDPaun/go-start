package repositories

import (
	"fmt"

	"github.com/MDPaun/go-start/tree/main/website/pkg/account/staff/models"
	postgres "github.com/MDPaun/go-start/tree/main/website/pkg/storage"
)

//Staff use from models
type Staff models.Staff

//StaffService use from models
type StaffRepository models.StaffRepository

//GetStaff get one Staff
func (s Staff) GetStaff() ([]*models.Staff, error) {

	db := postgres.ConnectDB()
	var staff []*models.Staff

	rows, err := db.Query(`SELECT id, email, fullname, image, status, date_added FROM staff ;`)
	check(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&s.ID, &s.Email, &s.FullName, &s.Image, &s.Status, &s.DateAdded)
		st := models.Staff(s)
		staff = append(staff, &st)
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, err
	}

	return staff, nil
	// var fullname string
	// for rows.Next() {
	// 	err = rows.Scan(&fullname)
	// 	check(err)
	// 	fmt.Println("RETRIEVED RECORD:", fullname)
	// }

	// return &fullname, nil

	// row := client.DbClient.QueryRow("select Staff.id,Staff.Staff_name,Staff.age,Staff.description,Staff.category_id,category.category_name,Staff.breed_id,breed.breed_name,Staff.location_id,location.location_name from Staff inner join category on Staff.category_id = category.id inner join breed on Staff.breed_id=breed.id inner join location on Staff.location_id = location.id where Staff.id=$1;", id)
	// switch err := row.Scan(&p.ID, &p.Name, &p.Age, &p.Description, &p.StaffCategory.ID, &p.StaffCategory.CategoryName, &p.StaffBreed.ID, &p.StaffBreed.BreedName, &p.StaffLocation.ID, &p.StaffLocation.LocationName); err {
	// case sql.ErrNoRows:
	// 	return nil, errors.New("No matching records")
	// case nil:
	// 	category := models.Staff(p)
	// 	return &category, nil
	// default:
	// 	return nil, err
	// }
}

// NewStaffRepo ..
// func GetStaff(db *sql.DB) {

// 	row := db.QueryRow("SELECT * FROM staff")
// 	fmt.Println(row)
// 	// return row

// }

// // FindByID ..
// func (r *StaffRepo) FindByID(ID int) (*models.Staff, error) {
// 	return &models.Staff{}, nil
// }

// // Save ..
// func (r *StaffRepo) Save(staff *models.Staff) error {
// 	return nil
// }
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
