package main

import (
	"github.com/danielargoe/winyp/internal"
)

func main() {
	key := internal.GetDotEnv(".env")
	loc := internal.GetLOCData()
	emoji := internal.GetIPData()
	data := internal.GetData(key, loc)
	s, m, l, f := internal.GetFlags()

	internal.GetWeather(s, m, l, f, data, emoji)
}
