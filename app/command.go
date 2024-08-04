package main

import (
	"bytes"
	"compress/gzip"
	"os"
	"strings"
)

func handleEcho(req request) response {
	msg := strings.TrimPrefix(req.path, "/echo/")
	if !contains(req.acceptEncoding, "gzip") {
		return newSuccessResponseWithBody(req, "text/plain", msg)
	}
	return newSuccessResponseWithBody(req, "text/plain", gzipString(msg))
}

func handleUserAgent(req request) response {
	return newSuccessResponseWithBody(req, "text/plain", req.userAgent)
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
	return newSuccessResponseWithBody(req, "application/octet-stream", string(data))
}

func handleCreateFile(req request) response {
	path := strings.TrimPrefix(req.path, "/files/")
	err := os.WriteFile(_metaInfo.basePath+path, []byte(req.body), 0644)
	if err != nil {
		return new404Response()
	}
	return newCreatedResponse()
}

func gzipString(input string) string {
	var compressedData bytes.Buffer
	gz := gzip.NewWriter(&compressedData)

	_, err := gz.Write([]byte(input))
	if err != nil {
		os.Exit(-1)
	}

	err = gz.Close()
	if err != nil {
		os.Exit(-1)
	}

	return compressedData.String()
}
