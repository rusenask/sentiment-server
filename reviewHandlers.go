package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleGetProductReview(r http.ResponseWriter, req *http.Request) {
	r.Header().Add("Content-Type", "application/json")

	if req.ContentLength < 1 {
		r.WriteHeader(http.StatusBadRequest)
		r.Write([]byte(fmt.Sprintf(`{"message": "no text passed. Cannot run review sentiment analysis"}`)))
		log.Printf("POST /review > ERROR: no text passed\n")
		return
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil && err != io.EOF {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(fmt.Sprintf(`{"message": "ERROR: error reading request body", "error": "%v"}`, err.Error())))
		log.Printf("POST /review > ERROR: couldn't read request body\n\t%v\n", err)
		return
	}

	j := ReviewJSON{}
	err = json.Unmarshal(data, &j)
	if err != nil {
		r.WriteHeader(http.StatusBadRequest)
		r.Write([]byte(fmt.Sprintf(`{"message": "ERROR: error unmarshalling given JSON into expected format", "error": "%v"}`, err.Error())))
		log.Printf("POST /analyze > ERROR: error unmarshalling given JSON\n\t%v\n", err)
		return
	}

	r.WriteHeader(http.StatusOK)
	r.Write([]byte("Not implemented yet"))

	count++
	log.Printf("POST /review \n")
}
