package internal

import (
	"flag"
	"fmt"
	"os"
)

func GetFlags() (bool, bool, bool, bool) {
	s := flag.Bool("s", false, "short")
	m := flag.Bool("m", false, "medium")
	l := flag.Bool("l", false, "long")
	f := flag.Bool("f", false, "full")
	flag.Parse()

	if !(*s && !*m && !*l && !*f) && !(!*s && *m && !*l && !*f) && !(!*s && !*m && *l && !*f) && !(!*s && !*m && !*l && *f) {
		LogFlagsError()
	}

	return *s, *m, *l, *f
}

func LogFlagsError() {
	fmt.Printf("Usage of %s: -s <short> -m <medium> -l <long> -f <full> \n", os.Args[0])
	os.Exit(1)
}
