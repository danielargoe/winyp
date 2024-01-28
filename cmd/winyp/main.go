package main

import (
	"fmt"

	"github.com/danielargoe/winyp/internal"
)

func main() {
	key := internal.GetDotEnv()
	loc, emoji := internal.GetIPData()
	data := internal.GetData(key, loc)

	fmt.Println(emoji)
	fmt.Println(data.Location.Name)
	fmt.Println(data.Location.Country)
	fmt.Println(data.Current.Temp_f)
}


