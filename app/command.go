package main

import "strings"

func handleEcho(str string) response {
	msg := strings.TrimPrefix(str, "/echo/")
	return newSuccessResponseWithBody(msg)
}
