package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/anishbista60/Golang/book-store/pkg/utils"
	"github.com/anishbista60/Golang/book-store/pkg/models"
	"github.com/gorilla/mux"
)
	
var newBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)

    if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
    parameters := mux.Vars(r)
    bookid, ok := parameters["bookid"]

    if !ok {
        http.Error(w, "Missing bookid parameter", http.StatusBadRequest)
        return
    }

    id, err := strconv.ParseInt(bookid, 10, 64)
    if err != nil {
        http.Error(w, "Invalid bookid parameter", http.StatusBadRequest)
        return
    }

    bookDetails,_:= models.GetBookById(id)
  

    res, err := json.Marshal(bookDetails)
    if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}


func CreateBook(w http.ResponseWriter, r*http.Request){
	createbook := &models.Book{}
	utils.ParseBody(r, createbook)
	b:= createbook.CreateBook()

	res, err := json.Marshal(b)
    if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    parameter := mux.Vars(r)
    bookid, ok := parameter["bookid"]

    if !ok {
        http.Error(w, "Missing bookid parameter", http.StatusBadRequest)
        return
    }

    id, err := strconv.ParseInt(bookid, 10, 64)
    if err != nil {
        http.Error(w, "Invalid bookid parameter", http.StatusBadRequest)
        return
    }

    book := models.DeleteBook(id)
    res, err := json.Marshal(book)
    if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook  = &models.Book{}
	utils.ParseBody(r, updateBook)
	parameter:= mux.Vars(r)

	bookid := parameter["bookid"]
	id,err := strconv.ParseInt(bookid,0,0)
	if err != nil{
		fmt.Println("Error while parsing")
	}
	booksDetails ,db := models.GetBookById(id)

	if updateBook.Name != ""{
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		booksDetails.Publication = updateBook.Publication
	}
	db.Save(&booksDetails)
	res, err := json.Marshal(booksDetails)
    if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
