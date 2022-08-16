package main

import (
	"encoding/json"
	"github.com/catena-x/gh-org-checks/pkg/data"
	"github.com/catena-x/gh-org-checks/pkg/testers"
	"github.com/catena-x/gh-org-checks/pkg/testrunner"
	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	orgReport  = data.OrgReports{}
	testRunner *testrunner.TestRunner
)

func main() {

	log.Printf("Starting service ...")
	setLogLevel()

	testRunner = testrunner.NewTestRunner()
	testRunner.AddToTestSuites(testers.NewReadMeTester)
	scheduleCronJobs()

	router := mux.NewRouter()
	router.HandleFunc("/report", returnOrgReport).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", router))

}

func setLogLevel() {
	log.SetLevel(log.DebugLevel)
}

func scheduleCronJobs() {
	log.Println("scheduled test cronjob")
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().Do(func() {
		go updateTestReport()
	},
	)

	s.StartAsync()
}

func updateTestReport() {
	log.Println("update test report")
	orgReport = testRunner.PerformRepoChecks()
}

func returnOrgReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("returning test report")
	json.NewEncoder(w).Encode(orgReport)
}
