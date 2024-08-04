package main

import "strings"

func handleEcho(str string) response {
	msg := strings.TrimPrefix(str, "/echo/")
	return newSuccessResponseWithBody(msg)
}

func handleUserAgent(req request) response {
	return newSuccessResponseWithBody(req.userAgent)
}
