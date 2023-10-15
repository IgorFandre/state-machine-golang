package state_machine

import (
	"fmt"
	"log"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getSliceString(slice []int) string {
	delim := ", "
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), delim), "[]")
}
