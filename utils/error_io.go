package util

import "log"

func RaiseError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}