package main

import (
	"github.com/cantillo16/bia_energy/src/connection"
	"github.com/cantillo16/bia_energy/src/importer"
)

func main() {

	importer.ImportCSVToDB(connection.Connect(), "src/importer/seeds/test_bia.csv")

}
