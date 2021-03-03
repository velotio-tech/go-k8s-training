package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

type RootHandler struct {
	PathMapping map[string]func(http.ResponseWriter, *http.Request)
}

func (handler *RootHandler) Init() {
	handler.PathMapping = make(map[string]func(http.ResponseWriter, *http.Request))
}

func (handler RootHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	pathChunks := strings.Split(strings.TrimSpace(req.URL.Path), "/")

	handlerFunc, present := handler.PathMapping[pathChunks[1]]

	if !present {
		fmt.Println("Path", req.URL.Path, "not found")
		http.NotFound(resp, req)
		return
	}

	handlerFunc(resp, req)
}
