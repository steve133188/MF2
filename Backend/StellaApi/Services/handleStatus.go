package Services

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
		w.Header().Set("Context-Type", "application/json")

		w.Write([]byte("failed"))
	}
	return
}

func HandleResult(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
	fmt.Println(bytes.NewBuffer(body))
}
