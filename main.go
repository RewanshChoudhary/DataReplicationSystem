package main

import (
	"github.com/RewanshChoudhary/DataReplicationSystem/DatabaseFunctions"
	"github.com/RewanshChoudhary/DataReplicationSystem/util"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	util.HandleError(err)

	// DEBUG: Check variables
	DatabaseFunctions.Init_Db()

}
