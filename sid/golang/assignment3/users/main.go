package main

import "github.com/farkaskid/go-k8s-training/assignment3/users/server"

func main() {
	server.New(2224).ListenAndServe()
}
