package main

import (
	"fmt"
	"strings"
)

type response struct {
	version    string
	statusCode int
	reason     string

	contentType   string
	contentLength int

	body string
}

func newSuccessResponse() response {
	return response{
		version:    "HTTP/1.1",
		statusCode: 200,
		reason:     "OK",
	}
}

func new404Response() response {
	return response{
		version:    "HTTP/1.1",
		statusCode: 404,
		reason:     "Not Found",
	}
}

func newSuccessResponseWithBody(contentType, body string) response {
	return response{
		version:       "HTTP/1.1",
		statusCode:    200,
		reason:        "OK",
		contentType:   contentType,
		contentLength: len(body),
		body:          body,
	}
}

func (r response) toByte() []byte {
	// header
	str := strings.Join([]string{
		r.version, fmt.Sprintf("%d", r.statusCode), r.reason,
	}, " ")
	str += "\r\n"

	header := ""
	if r.contentType != "" {
		header += fmt.Sprintf("Content-Type: %s\r\n", r.contentType)
	}
	if r.contentLength != 0 {
		header += fmt.Sprintf("Content-Length: %d\r\n", r.contentLength)
	}
	if len(header) > 0 {
		str += header + "\r\n"
	} else {
		str += "\r\n"
	}

	if len(r.body) > 0 {
		str += r.body
	}

	return []byte(str)
}
