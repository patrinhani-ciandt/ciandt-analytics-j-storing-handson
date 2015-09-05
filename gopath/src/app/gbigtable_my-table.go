package main

import (
	"fmt"
	// "time"
	// "encoding/json"

	// "google.golang.org/cloud/bigtable"

	"ciandt.golang.org/libs/gbigtable"
	// "ciandt.golang.org/libs/ioutil"
)

const (
	KeyJsonFilePath = "/home/key.json"
)

func main() {

	var bigtableClientConnData = gbigtable.ClientConnectionData {
		Project: "bigdatagarage",
		Zone: "us-central1-c",
		Cluster: "workshopanalytics",
		KeyJsonFilePath: KeyJsonFilePath,
	}

	createTable_my_table(bigtableClientConnData);

	importData_from_csv_into_my_table(bigtableClientConnData);

	insertData_on_my_table(bigtableClientConnData);

	readData_from_my_table(bigtableClientConnData);
}

func createTable_my_table(bigtableClientConnData gbigtable.ClientConnectionData) {
	
	fmt.Println("Creating my-table ...");
}

func importData_from_csv_into_my_table(bigtableClientConnData gbigtable.ClientConnectionData) {

	fmt.Println("Importing data from csv into my-table ...");
}

func insertData_on_my_table(bigtableClientConnData gbigtable.ClientConnectionData) {

	fmt.Println("Inserting data on my-table ...");
}

func readData_from_my_table(bigtableClientConnData gbigtable.ClientConnectionData) {

	fmt.Println("Reading data from my-table ...");
}