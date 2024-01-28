package internal

import (
	"log"

	"github.com/ipinfo/go/v2/ipinfo"
)

func GetIPData() (string, string) {
	loc, err := ipinfo.GetIPLocation(nil)
	if err != nil {
		log.Fatal(err)
	}

	country, err := ipinfo.GetIPCountry(nil)
	if err != nil {
		log.Fatal(err)
	}

	emoji := ipinfo.GetCountryFlagEmoji(country)

	return loc, emoji
}