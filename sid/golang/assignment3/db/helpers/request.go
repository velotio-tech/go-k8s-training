package helpers

import (
	"net/http"
	"strconv"
)

func GetIDs(req *http.Request) ([]int, error) {
	queryParams := req.URL.Query()

	ids, ok := queryParams["ids"]
	var idInts []int

	if !ok {
		return idInts, nil
	} else {
		idInts = make([]int, len(ids))
		for index, id := range ids {
			idInt, err := strconv.Atoi(id)

			if err != nil {
				return idInts, err
			}

			idInts[index] = idInt
		}
		return idInts, nil
	}
}
