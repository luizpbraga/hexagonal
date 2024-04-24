package main

import (
	"log"
	"os"
	"time"

	"github.com/luizpbraga/hexagonal/adapters/db"
	gRPC "github.com/luizpbraga/hexagonal/adapters/grpc"
	"github.com/luizpbraga/hexagonal/application/api"
	"github.com/luizpbraga/hexagonal/application/arithmetic"
)

func main() {
	time.Sleep(time.Second * 30)
	dsn := os.Getenv("DSN")
	dbDriver := os.Getenv("DB_DRIVER")
	dbAdapter, err := db.NewAdapter(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer dbAdapter.Close()

	log.Println("DB CONNECTED")

	core := arithmetic.New()

	applicationAPI := api.NewApplication(dbAdapter, core)

	log.Println("APPLICATION ON")

	grpcAdapter := gRPC.NewGRPCService(applicationAPI)

	log.Println("RUNNING GRPC SERVER")

	grpcAdapter.Run()
}
