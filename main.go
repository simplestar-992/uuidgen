package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
)

var (
	count  = flag.Int("c", 1, "Number of UUIDs")
	upper  = flag.Bool("u", false, "Uppercase")
)

func main() {
	flag.Parse()
	for i := 0; i < *count; i++ {
		id := uuid.New()
		s := id.String()
		if *upper { s = fmt.Sprintf("%X", id) }
		fmt.Println(s)
	}
}
