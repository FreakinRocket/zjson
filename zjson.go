package zjson

// ### CONSTANTS ###
import (
	"encoding/json"
	"log"
	"os"
)

// loads config into config struct from given file name
func LoadJSON(objectPointer *any, jsonFilePath string) {
	jsonFile, err := os.Open(jsonFilePath)
	ChkError(err)
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(objectPointer)
	ChkError(err)
}

// saves config to file from struct
func SaveJSON(object any, jsonFilePath string) {
	jsonData, err := json.MarshalIndent(object, "", " ")
	ChkError(err)
	err = os.WriteFile(jsonFilePath, jsonData, 0644)
	ChkError(err)
}

// saves programming time and log.fatalLn on an error
func ChkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
