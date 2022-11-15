package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type CssColorSpec struct {
	Specification string `json:"specification"`
	Keyword       string `json:"keyword"`
	RgbHexValue   string `json:"rgb_hex_value"`
}

/**
* Read the CSV file and return output as JSON
 */
func main() {
	url := "https://gist.githubusercontent.com/jr-frazier/1feabb463768158e788c5247f37fd1a4/raw/csv-color-spec.csv"

	// GET request
	// http.Get("https://www.google.com")
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("error: can't call %s", url)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Convert the response body from csv to json
	// csvToJson(body)
	csvReader := csv.NewReader(resp.Body)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var spec CssColorSpec
	var specList []CssColorSpec

	for _, row := range data {
		spec.Specification = row[0]
		spec.Keyword = row[1]
		spec.RgbHexValue = row[2]
		specList = append(specList, spec)
	}

	// Convert to JSON
	jsonData, err := json.Marshal(specList)
	if err != nil {
		log.Fatal(err)
	}

	// Print the JSON
	log.Println(string(jsonData))

}
