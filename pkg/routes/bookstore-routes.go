package routes

import (
	"github.com/YashBubna/library_management/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{bookID}", controllers.GetBookByID).Methods("GET")
	r.HandleFunc("/book/{bookID}", controllers.UpdateBookByID).Methods("POST")
	r.HandleFunc("/book/{bookID}", controllers.DeleteBookByID).Methods("DELETE")
}
