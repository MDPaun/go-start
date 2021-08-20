package main

import (
	"net/http"

	"github.com/MDPaun/go-start/tree/main/website/cmd/staff"
)

// var (
// 	router = echo.New()
// )

// func d(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "dog dog dog")
// }

// func c(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "cat cat cat")
// }

func main() {
	mapUrls()
	http.ListenAndServe(":8080", nil)
}

func mapUrls() {
	staff.StaffUrls()

}

// func main() {
// 	mapUrls()
// 	router.Logger.Fatal(router.Start(":3000"))
// }

// func mapUrls() {
// 	router.POST("/category", rest.CreateCategory)
// 	router.GET("/category", rest.GetAllCategory)
// 	router.GET("/category/:id", rest.GetCategory)
// 	router.POST("/breed", rest.CreateBreed)
// 	router.GET("/breed/:category_id", rest.GetBreedByCategory)
// 	router.POST("/location", rest.CreateLocation)
// 	router.GET("/location", rest.GetAllLocation)
// 	router.GET("/pet/:category_id", rest.GetPetByCategory)
// 	router.POST("/pet", rest.CreatePet)
// 	router.GET("/pet", rest.GetPet)
// 	router.DELETE("/pet", rest.DeletePet)
// }
