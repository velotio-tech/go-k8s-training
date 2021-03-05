package main

import "github.com/farkaskid/go-k8s-training/assignment3/orders/server"

func main() {
	server.New(2223).ListenAndServe()
}
