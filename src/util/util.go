package util

import (
	"strconv"
	"path/filepath"
	"log"
	"strings"
)

func HostPortToAddress(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}

func UrlToHost(url string) string {
	return strings.Split(url, ":")[0]
}

func AbsolutePath(relpath string) string {
	absolutePath, err := filepath.Abs(relpath)
	if err != nil {
		log.Println("current path error:", err)
	}
	return absolutePath
}

func HomePath() string {
	return AbsolutePath(".")
}