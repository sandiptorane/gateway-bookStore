package server

import (
	"context"
	"fmt"
	books "gateway-bookstore/bookproto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Run runs the gRPC-Gateway, dialling the provided address.
func Run(dialAddr string) error {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	gwmux := runtime.NewServeMux()
	//register Book Service handler to gwmux
	err = books.RegisterBookServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	gatewayAddr := ":8082"  //address for gateway rest api
	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/bookstore") {
				gwmux.ServeHTTP(w, r)
				return
			}
		}),
	}
	// return gmServer
	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())


}

func (s *Server)InitializeServer(){
	book1 := &books.Book{
		Id:      1,
		Title:   "clean code",
		Author:  "uncle bob",
		Content: "how to do clean code",
	}
	book2 :=&books.Book{
		Id:      2,
		Title:   "C Programiming",
		Author:  "kanetkar",
		Content: "basic and intermediate course",
	}
	var bookStore []*books.Book
	bookStore = append(bookStore,book1,book2)
	s.bookStore = bookStore
}
