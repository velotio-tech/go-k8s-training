package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func writeJSON(w http.ResponseWriter, src any) {
	marshalledResponse, err := json.MarshalIndent(src, "", "    ")
	if err != nil {
		log.Println("marshalling err ", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "marshalling ", err)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledResponse)
}
