package main

import (
	"io"
	"log"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("postgres").Load()
	if err != nil {
		log.Fatal("failed to load gorm schema: ", err)
	}
	io.WriteString(os.Stdout, stmts)
}
