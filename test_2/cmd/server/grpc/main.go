package main

import (
	"log"
	"net"

	pb "github.com/inouttt/stkt/test_2"
	"github.com/inouttt/stkt/test_2/internal/config"
	ll "github.com/inouttt/stkt/test_2/internal/log"
	"github.com/inouttt/stkt/test_2/internal/movie"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	configuration := config.New()
	args := configuration.SetArgs()
	dbConf := config.GetDbConf(args)

	mysqlCon, err := config.NewMysqlSession(dbConf)
	if err != nil {
		log.Panicln("Failed initialize DB")
	}

	// init repo
	logRepo := ll.NewRepository(mysqlCon)

	// init services
	logger := ll.Newservices(logRepo)
	ms := movie.NewServices(args.UrlTarget, logger)

	// init grpcs service
	mgs := movie.NewGrpc(ms)

	// Open socket
	lis, err := net.Listen("tcp", ":"+args.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMovieServiceServer(s, mgs)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
