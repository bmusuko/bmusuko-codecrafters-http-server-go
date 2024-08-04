package main

import (
	"fmt"
	"strings"
)

type request struct {
	method    string
	path      string
	version   string
	host      string
	userAgent string
	accept    string
}

func newRequest(rawStr string) request {
	strs := strings.Split(rawStr, "\r\n")
	r := request{}
	r.setRequestLine(strs[0])
	r.setHeader(strs[1:])
	return r
}

func (r *request) setRequestLine(str string) {
	parts := strings.Split(str, " ")
	r.method = parts[0]
	r.path = parts[1]
	r.version = parts[2]
}

func (r *request) setHeader(strs []string) {
	for _, str := range strs {
		if str == "" {
			continue
		}

		parts := strings.Split(str, ": ")
		key := parts[0]
		switch key {
		case "Host":
			r.host = parts[1]
		case "User-Agent":
			r.userAgent = parts[1]
		case "Accept":
			r.accept = parts[1]
		default:
			fmt.Printf("unknown key val %s\n", str)
		}
	}
}
