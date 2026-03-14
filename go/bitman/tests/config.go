package tests

import (
	"os"
	"encoding/json"
)

type TestFields struct {
	Value int `json:"Value"`
	Expected int `json:"Expected"`
}

type TestsList struct {
	Test string
	Cases  []TestFields
}

var Cases TestsList

func LoadTests(filename string) error {
	fl, err := os.Open(filename)
	defer fl.Close()

	if err != nil {
		return err
	}

	err = json.NewDecoder(fl).Decode(&Cases)

	if err != nil {
		return err
	}
	return nil
}
