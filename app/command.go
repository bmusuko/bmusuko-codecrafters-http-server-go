package main

import "strings"

func handleEcho(req request) response {
	msg := strings.TrimPrefix(req.path, "/echo/")
	return newSuccessResponseWithBody(msg)
}

func handleUserAgent(req request) response {
	return newSuccessResponseWithBody(req.userAgent)
}
