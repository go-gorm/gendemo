package util

import "log"

func CatchErr(logStr string, err error) {
	if err != nil {
		log.Fatal(logStr+": %s", err)
	}
}
