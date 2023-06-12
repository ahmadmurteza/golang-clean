package main

import (
	"fmt"
	"golang-clean/internal/adapters/app/api"
	"golang-clean/internal/adapters/core/arithmetic"
	"golang-clean/internal/adapters/framework/right/db"
	"golang-clean/internal/ports"
	"log"
	"os"
)

func main() {
	var err error

	// ports
	var dbasAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort

	// database connection
	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbasAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("db connection error:", err)
	}
	defer dbasAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbasAdapter, core)

	fmt.Println(appAdapter)
}
