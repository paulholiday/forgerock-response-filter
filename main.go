package main

import (
	"bufio"
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

	if (os.Args[3] != "") {
		
		//get list of users to deactivate
		activeUsersFile, err := os.Open(os.Args[1])

		if err != nil {
			log.Fatal(err)
		}

		fileScanner := bufio.NewScanner(activeUsersFile)
		fileScanner.Split(bufio.ScanLines)
		var activeUsers []string

		for fileScanner.Scan() {
			activeUsers = append(activeUsers, fileScanner.Text())
		}

		inactiveForgeRockUsersFile, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		fileScanner = bufio.NewScanner(inactiveForgeRockUsersFile)
		fileScanner.Split(bufio.ScanLines)
		var inactiveFRUsers []string

		for fileScanner.Scan() {
			inactiveFRUsers = append(inactiveFRUsers, fileScanner.Text())
		}

		for i, activeUser := range activeUsers {
			for j, inactiveFRUser := range inactiveFRUsers {
				if (inactiveFRUser == activeUser) {
					//remove
					inactiveFRUsers = remove(inactiveFRUsers, j)
					fmt.Printf("Removed %d\n", i)
					break
				}
			}
		}

		for _, inactiveFRUser := range inactiveFRUsers {
			fmt.Printf("%s\n", inactiveFRUser)
		}
		
	} else if (os.Args[2] != "") {
		compareFiles()
	} else {
		//filter results
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
	
}

func compareFiles() {

	firstFile, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(firstFile)
	fileScanner.Split(bufio.ScanLines)
	var firstFileLines []string

	for fileScanner.Scan() {
		firstFileLines = append(firstFileLines, fileScanner.Text())
	}

	secondFile, err := os.Open(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	fileScanner = bufio.NewScanner(secondFile)
	fileScanner.Split(bufio.ScanLines)
	var secondFileLines []string

	for fileScanner.Scan() {
		secondFileLines = append(secondFileLines, fileScanner.Text())
	}

	//output common users in both files
	for _, fileLine := range firstFileLines {
		if contains(secondFileLines, fileLine) {
			fmt.Printf("%s\n", fileLine)
		}
	}
}

// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func remove(s []string, i int) []string {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}
