package main

import (
	"log"

	"github.com/luizpbraga/hexagonal/adapters/db"
	gRPC "github.com/luizpbraga/hexagonal/adapters/grpc"
	"github.com/luizpbraga/hexagonal/application/api"
	"github.com/luizpbraga/hexagonal/application/arithmetic"
)

func main() {
	dbAdapter, err := db.NewAdapter("mysql", "test:test@tcp(localhost)/hexagonal")
	if err != nil {
		log.Fatal(err)
	}
	defer dbAdapter.Close()

	log.Println("DB CONNECTED")

	core := arithmetic.New()

	applicationAPI := api.NewApplication(dbAdapter, core)

	log.Println("APPLICATION ON")

	grpcAdapter := gRPC.NewGRPCService(applicationAPI)

	log.Println("RUNNING GRPC SERVER :9000")

	grpcAdapter.Run()
}
