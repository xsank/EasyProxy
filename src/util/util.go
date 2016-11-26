package util

import (
	"strconv"
	"path/filepath"
	"log"
	"strings"
	"reflect"
	"net/http"
	_ "net/http/pprof"
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

func SliceIndex(slice interface{}, element interface{}) int {
	index := -1
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		return index
	}
	ev := reflect.ValueOf(element).Interface()
	length := sv.Len()
	for i := 0; i < length; i++ {
		iv := sv.Index(i).Interface()
		if reflect.DeepEqual(iv, ev) {
			index = i;
			break
		}
	}
	return index
}

func Debug(open bool) {
	if open {
		go func() {
			http.ListenAndServe("localhost:6060", nil)
		}()
		log.Println("curl http://localhost:6060/debug/pprof/ to get more information")
	}
}