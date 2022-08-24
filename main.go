package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Profile struct {
	Oid string `json:"oid"`
}

func main() {
	fmt.Println("start to compare json files...")
	prodRawProfiles, err := os.Open("prod-profiles.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("successfully opened prod-profiles.json")
	defer prodRawProfiles.Close()

	prodProfilesByte, _ := io.ReadAll(prodRawProfiles)
	var prodProfiles []Profile
	err = json.Unmarshal(prodProfilesByte, &prodProfiles)
	if err != nil {
		panic(err)
	}

	testRawProfiles, err := os.Open("test-profiles.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully opened prod-profiles.json")
	defer prodRawProfiles.Close()

	testProfilesByte, _ := io.ReadAll(testRawProfiles)
	var testProfiles []Profile
	err = json.Unmarshal(testProfilesByte, &testProfiles)
	if err != nil {
		panic(err)
	}

	// put prod file into map to enhance the efficiency
	pp := make(map[string]struct{}, len(prodProfiles))
	for _, profile := range prodProfiles {
		pp[profile.Oid] = struct{}{}
	}

	var diff []string
	for _, profile := range testProfiles {
		if _, found := pp[profile.Oid]; !found {
			diff = append(diff, profile.Oid)
		}
	}

	fmt.Println(len(diff))

	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for _, d := range diff {
		if _, err = f.WriteString(d + "\n"); err != nil {
			panic(err)
		}
	}
	fmt.Println("end to compare json files...")
}
