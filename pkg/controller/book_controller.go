package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Teplotax/study-go-bookstore/pkg/model"
	"github.com/Teplotax/study-go-bookstore/pkg/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//var NewBook model.Book

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	books := model.GetAllBooks()
	response, _ := json.Marshal(books)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(response)
	if err != nil {
		return
	}
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}
	book, _ := model.GetBookById(id)
	response, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(response)
	if err != nil {
		return
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	createBook := &model.Book{}
	util.ParseBody(request, createBook)
	book := createBook.CreateBook()
	response, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	_, err := writer.Write(response)
	if err != nil {
		return
	}
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}
	book := model.DeleteBook(Id)
	_, _ = json.Marshal(book)
	writer.Header().Set("Content-Type", "application.json")
	writer.WriteHeader(http.StatusNoContent)
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	updateBook := &model.Book{}
	util.ParseBody(request, updateBook)
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}

	book, db := model.GetBookById(id)
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	db.Save(&book)
	response, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(response)
	if err != nil {
		return
	}
}
