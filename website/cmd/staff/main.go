package staff

import (
	"net/http"

	"github.com/MDPaun/go-start/tree/main/website/internal/staff/rest"
)

func StaffUrls() {
	http.Handle("/dog", http.HandlerFunc(rest.GetStaff))
	http.Handle("/staff", http.HandlerFunc(rest.GetStaff))
}
