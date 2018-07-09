package command

import (
	"fmt"
	"os"
	"scm/lib/utils"
	"time"
)

/*
	file structure
	.scm
		branch/        <-- used to store branch information
			master.json
		commit/        <-- used to store commits
			commithash.json
		objects/       <-- used to store file objects
			fileObj.json
		config.json    <-- config of the repo
*/

//Init initialize an empty repo
func Init(args []string) {
	if err := os.Mkdir(".scm", 0777); err == nil {
		createInitFiles()
		fmt.Print("scm repo initialized.")
	} else {
		fmt.Println("an error occured while initializing scm repo.")
		fmt.Println("the repo may already be initialized.")
		fmt.Println(err)
	}
}

func createInitFiles() {
	if err := os.Mkdir("./.scm/branch", 0777); err != nil {
		fmt.Println("an error occured while initializing scm repo.")
		fmt.Println(err)
	}
	if err := os.Mkdir("./.scm/commit", 0777); err != nil {
		fmt.Println("an error occured while initializing scm repo.")
		fmt.Println(err)
	}
	if err := os.Mkdir("./.scm/objects", 0777); err != nil {
		fmt.Println("an error occured while initializing scm repo.")
		fmt.Println(err)
	}
	initConfig()
}

type configStruct struct {
	name         string
	creationTime time.Time
}

func initConfig() {
	config := configStruct{
		"jake",
		time.Now(),
	}
	utils.StructToJSON("./.scm/config.json", config)
	test := new(configStruct)
	utils.JSONToStruct("./.scm/config.json", &test)
	fmt.Println(test)
}
