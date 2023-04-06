package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type APIResponse struct {
	Result                  []Data `json:"result"`
	ResultCount             int32  `json:"resultCount"`
	PagedResultsCookie      string `json:"pagedResultsCookie"`
	TotalPagedResultsPolicy string `json:"totalPagedResultsPolicy"`
	TotalPagedResults       int32  `json:"totalPagedResults"`
	RemainingPagedResults   int32  `json:"remainingPagedResults"`
}

type Data struct {
	Id        string   `json:"_id"`
	Rev       string   `json:"_rev"`
	UserName  string   `json:"userName"`
}

func main() {

	files, _ := os.ReadDir(os.Args[1])

	for _, file := range files {
		content, err := os.ReadFile(os.Args[1] + "/" + file.Name())

		if err != nil {
			log.Fatal(err)
		}

		data := APIResponse{}
		json.Unmarshal([]byte(content), &data)

		for _, resultData := range data.Result {
			fmt.Printf("%s\n", resultData.UserName)
		}
	}

}
