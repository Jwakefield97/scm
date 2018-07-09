package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//TODO: fix not serializing data correctly
//StructToJSON save a struct as json at a specified location
func StructToJSON(filePath string, object interface{}) {
	jsonData, err := json.Marshal(object)
	if err != nil {
		fmt.Println("an error occured while marshalling struct")
		fmt.Println(err)
	} else {
		if err := ioutil.WriteFile(filePath, jsonData, 0644); err != nil {
			fmt.Println("an error occured while writing marshalled data")
			fmt.Println(err)
		}
	}
}

//JSONToStruct deserialize json file to a struct
func JSONToStruct(filePath string, returnObj interface{}) {
	if bytes, err := ioutil.ReadFile(filePath); err != nil {
		json.Unmarshal(bytes, returnObj)
	} else {
		fmt.Println("there was an error while unmarshalling data")
		fmt.Println(err)
	}
}
