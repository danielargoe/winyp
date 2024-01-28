package internal

import (
	"log"

	"github.com/ipinfo/go/v2/ipinfo"
)

func GetLOCData() string {
	loc, err := ipinfo.GetIPLocation(nil)
	if err != nil {
		log.Fatal(err)
	}

	return loc
}

func GetIPData() string {
	country, err := ipinfo.GetIPCountry(nil)
	if err != nil {
		log.Fatal(err)
	}

	emoji := ipinfo.GetCountryFlagEmoji(country)

	return emoji
}