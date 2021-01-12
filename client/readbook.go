package main

import (
	"context"
	"fmt"
	books "gateway-bookstore/bookproto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main(){
	var conn *grpc.ClientConn
	//dial for server conn
	conn,err:= grpc.Dial(":8081",grpc.WithInsecure())
	if err!=nil{
		log.Fatal("did not connect : ",err)
	}
	defer conn.Close()

	//initialize bookService
	bookClient := books.NewBookServiceClient(conn)
	response,err :=bookClient.ReadBook(context.Background(),&books.ReadBookRequest{Id: 3})
	if err!=nil{
		fmt.Println(err)
		return
	}
	book := response.GetBook()
	fmt.Println("requested book from bookstore:",book)

}
