package main

import (
	"fmt"
	"log"
	"net"

	"github.com/14jasimmtp/auth-svc/pkg/config"
	"github.com/14jasimmtp/auth-svc/pkg/db"
	"github.com/14jasimmtp/auth-svc/pkg/helper"
	"github.com/14jasimmtp/auth-svc/pkg/pb"
	"github.com/14jasimmtp/auth-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.NewConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Connection(c.DB_URL)

	jwt := helper.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
