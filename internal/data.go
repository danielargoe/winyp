package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetData(key string, loc string) Data {
	url := "http://api.weatherapi.com/v1/current.json?key=" + key + "&q=" + loc

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var j Data
	err = json.Unmarshal(body, &j)
	if err != nil {
		log.Fatal(err)
	}

	return j
}
