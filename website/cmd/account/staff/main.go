package rest

import (
	"net/http"

	s "github.com/MDPaun/go-start/tree/main/website/pkg/account/staff/controllers"
)

func StaffUrls() {
	http.HandleFunc("/dashboard", s.GetStaff)
	// http.Handle("/staff", http.HandlerFunc(rest.GetStaff))
}
