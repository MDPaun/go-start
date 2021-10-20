package controllers

import (
	"fmt"
	"net/http"

	db "github.com/MDPaun/go-start/tree/main/website/pkg/account/staff/storage"
)

// var sr db.StaffRepository
var staff db.Staff

func GetStaff(res http.ResponseWriter, req *http.Request) {
	// sr = &staff
	rows := staff.GetStaff()

	var staff struct{}
	for rows.Next() {
		err = rows.Scan(&staff)
		check(err)
		fmt.Fprintln(res, "RETRIEVED RECORD:", staff)
	}
}

// func GetStaff1(res http.ResponseWriter, req *http.Request) {

// rows, err := db.Query(`SELECT fullname FROM staff ;`)
// check(err)
// defer rows.Close()

// var fullname string
// for rows.Next() {
// 	err = rows.Scan(&fullname)
// 	check(err)
// 	fmt.Fprintln(res, "RETRIEVED RECORD:", fullname)
// }
// }

// BaseHandler will hold everything that controller needs
// type BaseHandler struct {
// 	staffRepo models.StaffRepository
// }

// // NewBaseHandler returns a new BaseHandler
// func NewBaseHandler(staffRepo models.StaffRepository) *BaseHandler {
// 	return &BaseHandler{
// 		staffRepo: staffRepo,
// 	}
// }
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
