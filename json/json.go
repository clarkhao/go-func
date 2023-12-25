package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type JSONItems = map[string]interface{}

// ReadJson is a function that read JSON file from fileName and
// return the JSON items as map and error
func ReadJson(fileName string) (*JSONItems, error) {
	byteData, err := os.ReadFile(fileName)
	if err != nil {
		err = fmt.Errorf("ReadJson when Readfile %v", err)
		return nil, err
	}
	var jsonItems JSONItems
	err = json.Unmarshal(byteData, &jsonItems)
	if err != nil {
		err = fmt.Errorf("ReadJson when Unmarshal file data: %v", err)
		return nil, err
	}
	return &jsonItems, nil
}

// WriteJson is a function that write some json items into a file
// return error
func WriteJson(fileName string, items *JSONItems) error {
	jsonData, err := json.Marshal(*items)
	if err != nil {
		err = fmt.Errorf("WriteJson when marshal to json data with %v", err)
		return err
	}
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		err = fmt.Errorf("WriteFile when writing file with %v", err)
		return err
	}
	return nil
}

// AddItems is a function that add some json items into target file
// return the number of items added and error
func AddItems(targetFile string, extensible string) (int, error) {
	original, err := ReadJson(targetFile)
	if err != nil {
		err = fmt.Errorf("AddItems when getting items from target file: %v", err)
		return 0, err
	}
	extense, err := ReadJson(extensible)
	if err != nil {
		err = fmt.Errorf("AddItems when getting items from extensible file: %v", err)
		return 0, err
	}
	count := 0
	for key, value := range *extense {
		(*original)[key] = value
		count++
	}
	err = WriteJson(targetFile, original)
	if err != nil {
		err = fmt.Errorf("AddItems when writing json data into target file with %v", err)
		return 0, err
	}
	return count, nil
}
