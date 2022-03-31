package controllers

import(

	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/federicopfund/go-bookstore/pkg/utils"// importar paquetes que estan en el Stock
	"github.com/federicopfund/go-bookstore/pkg/models"
)
var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res, _ :=json.Marshal(newBooks)// Convertirlo a Json
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)// basicamennte dara un 200
	w.Write(res)// enviar algo por Frontend
}
// Obtener un libro por ID
func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _:= models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)// enviar Json como resultado
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)// estatus OK
	w.Write(res) // enviamos la respuesta
}
//Escritor de respuesta 
func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)// enviamos la respuesta ejem Posman
}
// actualizar el registro ya creado con c aplicaciones 
func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(ID)// una ves crado actualizara el estado de book.
	if updateBook.Name != ""{// si actualizar book no es igual a cadena vacia
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{// los mismo para el autor
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{// y con las publicaciones
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)// una vez actualizado lo guardamos en nuestra base de datos.
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}