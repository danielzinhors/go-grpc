package main

import (
	"database/sql"
	"net"

	"github.com/danielzinhors/go-grpc/internal/database"
	"github.com/danielzinhors/go-grpc/internal/pb"
	"github.com/danielzinhors/go-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	categoryDB := database.NewCategory(db)
	categoryServie := service.NewCategoryService(*categoryDB)
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryServie)
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
