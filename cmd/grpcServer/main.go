package main

import (
	"database/sql"
	"net"

	"github.com/DuMendonca/grpc-full-cycle/internal/database"
	"github.com/DuMendonca/grpc-full-cycle/internal/pb"
	"github.com/DuMendonca/grpc-full-cycle/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err) //Exception in Java
	}
	defer db.Close() //Depois da execução fecha a conexão

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()                                // Criação de um Servidor gRPC
	pb.RegisterCategoryServiceServer(grpcServer, categoryService) //Vincular o Servidor com o Service (Gerado pelo protoc)
	reflection.Register(grpcServer)                               // Registra a API no servidor

	lis, err := net.Listen("tcp", ":50051") //Cria a conexão entre client e server na porta padrão (50051) do gRPC
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil { // .Serve() vai retornar caso exista algum erro na conexão.
		panic(err)
	}
}
