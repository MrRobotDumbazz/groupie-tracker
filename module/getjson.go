package module

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetJson(url string, object interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.Unmarshal(body, &object)
	if err != nil {
		return err
	}
	return nil
}

func GetRelation() error{
	err := GetJson(artisturl, &Module)
	if err != nil {
		return err
	}
	err = GetJson(relationurl, &relation)
	if err != nil {
		return err 
	}
	for i, w := range relation.Index {
		Module[i].DatesLocations = w.DatesLocations
	}
	return nil 
}
