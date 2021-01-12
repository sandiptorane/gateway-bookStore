package main

import (
	"context"
	"fmt"
	books "gateway-bookstore/bookproto"
	"google.golang.org/grpc"
	"log"
)

func main(){
	var conn *grpc.ClientConn
	//dial server it will returns connection
	conn,err:= grpc.Dial(":8081",grpc.WithInsecure())
	if err!=nil{
		log.Fatal("did not connect : ",err)
	}

	//initialize BookService
	bookClient := books.NewBookServiceClient(conn)
	fmt.Println("bookclient:",bookClient)
	book := &books.AddBookRequest{
		Book: &books.Book{
			Title:         "clean code",
			Author:        "uncle bob",
			Content:       "how to do clean code",
		},
	}
	response,err :=bookClient.AddBook(context.Background(),book)
	if err!=nil{
		log.Fatal("error when calling AddBook:",err.Error())
	}
	addedBook := response.GetBook()

	fmt.Println("New book added to bookstore:",addedBook)
}
