package parser

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// Parse parses the XML file and stores the result.
func Parse(name string, value interface{}) error {
	file, err := os.Open(name)
	defer file.Close()

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	return xml.Unmarshal(data, value)
}
