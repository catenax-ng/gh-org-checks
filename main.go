package main

import (
	"encoding/json"
	"fmt"
	"github.com/catena-x/gh-org-checks/pkg"
	"github.com/catena-x/gh-org-checks/pkg/testers"
	"log"
)

func main() {

	log.Printf("Starting service ...")

	testRunner := pkg.NewTestRunner()
	testRunner.AddToTestSuites(testers.NewReadMeChecker)
	result := testRunner.PerformRepoChecks()

	res, err := PrettyStruct(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
