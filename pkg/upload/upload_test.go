package upload

import (
	"log"
	"net/http"
	"testing"
)

func TestUploadHandler(t *testing.T) {
	http.HandleFunc("/switch/range/upload", UploadHandler)
	if err := http.ListenAndServe(":81", nil); err != nil {
		log.Fatalf("error to start http server:%s", err.Error())
	}
}
