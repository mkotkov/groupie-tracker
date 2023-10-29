package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"groupieTracker/models"
)

func FetchData(apiURL string) ([]byte, error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ParseData(data []byte, target interface{}) error {
	err := json.Unmarshal(data, target)
	if err != nil {
		return err
	}
	return nil
}

func FormatRelations(relations []models.Relation) map[string][]string {
	formattedRelations := make(map[string][]string)
	for _, relation := range relations {
		for location, dates := range relation.DatesLocation {
			formattedRelations[location] = dates
		}
	}
	return formattedRelations
}
