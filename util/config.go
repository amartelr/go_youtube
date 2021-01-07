package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/amartelr/go_youtube/entity"
)

// RestErrorResponse any
type RestErrorResponse struct {
	Code            int
	Message         string
	ValidationError map[string]interface{}
}

func LoadParams() (param *entity.Parameters) {

	filename, err := os.Open("assets/Params.json")
	if err != nil {
		log.Fatal(err)
	}

	defer filename.Close()

	data, err := ioutil.ReadAll(filename)

	if err != nil {
		log.Fatal(err)
	}

	result := entity.Parameters{}

	jsonErr := json.Unmarshal(data, &result)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return &result

}
