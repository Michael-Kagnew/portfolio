package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Projects struct {
	PrjNums int      `json: "prjnums"`
	Desc    []string `json: "desc"`
}

func GetProjects() {
	file, err := ioutil.ReadFile("test.json")

	if err != nil {
		fmt.Println("Couldn't open file")
	}

	var arr Projects
	err = json.Unmarshal(file, &arr)

	if err != nil {
		fmt.Println("couldn't read into array")
	}

	fmt.Println(arr)

}
