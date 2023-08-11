package route

import "github.com/gorilla/mux"
import "github.com/Teplotax/study-go-bookstore/pkg/controller"

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/books/{bookId}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controller.DeleteBook).Methods("DELETE")
}
