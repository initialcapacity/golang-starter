package datacollector

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func MakeRequest(w Task) ([]byte, error) {
	get, err := http.Get(w.Url)
	if err != nil {
		return nil, err
	}
	log.Println(fmt.Sprintf("Collecting data from %s", w.Url))
	body, err := io.ReadAll(get.Body)
	return body, nil
}
