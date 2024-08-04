package main

import (
	"fmt"
	"strconv"
	"strings"
)

type response struct {
	req request

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

func newCreatedResponse() response {
	return response{
		version:    "HTTP/1.1",
		statusCode: 201,
		reason:     "Created",
	}
}

func new404Response() response {
	return response{
		version:    "HTTP/1.1",
		statusCode: 404,
		reason:     "Not Found",
	}
}

func newSuccessResponseWithBody(req request, contentType, body string) response {
	return response{
		req:           req,
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
	if r.req.acceptEncoding == "gzip" {
		header += fmt.Sprintf("Content-Encoding: %s\r\n", r.req.acceptEncoding)
	}

	if len(header) > 0 {
		str += header + "\r\n"
	} else {
		str += "\r\n"
	}

	if len(r.body) > 0 {
		str += r.body
	}

	fmt.Printf("sent response %s\b", strconv.Quote(str))

	return []byte(str)
}
