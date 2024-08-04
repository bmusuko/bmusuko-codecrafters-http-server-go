package main

import (
	"os"
	"strings"
)

func handleEcho(req request) response {
	msg := strings.TrimPrefix(req.path, "/echo/")
	return newSuccessResponseWithBody("text/plain", msg)
}

func handleUserAgent(req request) response {
	return newSuccessResponseWithBody("text/plain", req.userAgent)
}

func handleFile(req request) response {
	if req.method == "POST" {
		return handleCreateFile(req)
	}
	return handleGetFile(req)
}

func handleGetFile(req request) response {
	path := strings.TrimPrefix(req.path, "/files/")
	data, err := os.ReadFile(_metaInfo.basePath + path)
	if err != nil {
		return new404Response()
	}
	return newSuccessResponseWithBody("application/octet-stream", string(data))
}

func handleCreateFile(req request) response {
	path := strings.TrimPrefix(req.path, "/files/")
	err := os.WriteFile(_metaInfo.basePath+path, []byte(req.body), 0644)
	if err != nil {
		return new404Response()
	}
	return newCreatedResponse()
}
