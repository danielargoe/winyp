package internal

import "fmt"

func GetWeather(s bool, m bool, l bool, f bool, data Data, emoji string) {
	if s {
		GetShortWeather(data, emoji)
	}
	if m {
		GetMediumWeather(data, emoji)
	}
	if l {
		GetLongWeather(data, emoji)
	}
	if f {
		GetFullWeather(data, emoji)
	}
}

func GetShortWeather(data Data, emoji string) {
	city := data.Location.Name
	temp := data.Current.Temp_c
	condition := data.Current.Condition.Text

	fmt.Printf("%s  %s: %s %.1f°c\n", emoji, city, condition, temp)
}

func GetMediumWeather(data Data, emoji string) {}

func GetLongWeather(data Data, emoji string) {}

func GetFullWeather(data Data, emoji string) {}
