package server

import (
	"context"
	"fmt"
	books "gateway-bookstore/bookproto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
      bookStore []*books.Book
}

func NewServer() *Server{
	return &Server{}
}

//AddBook get input book from AddBookRequest with blank id and
//return AddBookResponse by adding book to bookstore and assigning id to newbook
func (s *Server) AddBook(ctx context.Context, request *books.AddBookRequest) (*books.AddBookResponse, error) {
	bookList := s.GetBookList()  //get booklist
	length := len(bookList)
	lastID := bookList[length-1].Id  //get last book id
	log.Println("Last book Id :",lastID)
	newBook :=&books.Book{
		Id:      lastID + 1,
		Title:   request.Book.GetTitle(),
		Author:  request.Book.GetAuthor(),
		Content: request.Book.GetContent(),
	}
	s.bookStore = append(s.bookStore, newBook)
	return &books.AddBookResponse{
		Book: newBook,
	},nil
}

//GetBookList return booklist of bookStore
func (s *Server)GetBookList() []*books.Book{
	return s.bookStore
}

//ReadBook get ReadBookRequest which contains book id to read
// and return book through ReadBookResponse
func (s *Server) ReadBook(ctx context.Context, request *books.ReadBookRequest) (*books.ReadBookResponse, error) {
	id :=request.GetId()
	var responseBook *books.Book
	for i,book := range s.bookStore{
		if s.bookStore[i].Id==id{
			responseBook = book
			log.Println("returning requested book from ReadBook")
			return &books.ReadBookResponse{Book: responseBook},nil
		}
	}

	return nil,status.Errorf(codes.NotFound,"requested book %d not found",id)
}


//UpdateBook get UpdateBookRequest and return Updated book through UpdateBookResponse
func (s *Server) UpdateBook(ctx context.Context, request *books.UpdateBookRequest) (*books.UpdateBookResponse, error) {
	bookToUpdate :=request.GetBook()   //get book from request
	id := bookToUpdate.GetId()
	//search book to update
	for i,_ := range s.bookStore{
		if s.bookStore[i].Id==id{
			//if book id found update its data
			update := &books.Book{
				Id:      bookToUpdate.GetId(),
				Title:   bookToUpdate.GetTitle(),
				Author:  bookToUpdate.GetAuthor(),
				Content: bookToUpdate.GetContent(),
			}
			s.bookStore[i] = update
			return &books.UpdateBookResponse{Book: s.bookStore[i]},nil
		}
	}
	return nil,status.Errorf(codes.NotFound,"book %d not found for update",id)
}


//DeleteBook takes book id to delete and return boolean success value through DeleteBookResponse
func (s *Server) DeleteBook(ctx context.Context, request *books.DeleteBookRequest) (*books.DeleteBookResponse, error) {
	id := request.GetId()
	//search book to update
	for i,_ := range s.bookStore{
		if s.bookStore[i].Id==id{
			//if book id found delete it
			s.bookStore = append(s.bookStore[:i],s.bookStore[i+1:]...)
			fmt.Printf("bookstore after delete book %d \n %v",id,s.bookStore)
			return &books.DeleteBookResponse{Success: true},nil
		}
	}
	return nil,status.Errorf(codes.NotFound,"book %d not found for delete",id)
}

//ListBooks returns all books from bookStore
func (s *Server) ListBooks(ctx context.Context, request *books.ListBookRequest) (*books.ListBookResponse, error) {
	if len(s.bookStore)!=0{
		return &books.ListBookResponse{Books: s.bookStore},nil
	}
	return nil,status.Errorf(codes.NotFound,"bookStore is empty\n")
}

