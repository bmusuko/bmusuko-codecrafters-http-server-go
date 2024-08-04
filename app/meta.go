package main

import "flag"

type metaInfo struct {
	basePath string
}

var (
	_metaInfo metaInfo
)

func initMeta() {
	flag.StringVar(&_metaInfo.basePath, "directory", "", "Directory")
	flag.Parse()
}
