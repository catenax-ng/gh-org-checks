package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/catena-x/gh-org-checks/pkg/data"
	"github.com/catena-x/gh-org-checks/pkg/testers"
	"github.com/catena-x/gh-org-checks/pkg/testrunner"
	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	orgReport  = data.OrgReports{}
	testRunner *testrunner.TestRunner
)

func main() {

	log.Printf("Starting service ...")
	setLogLevel()

	initTestSuiteAndSchedule(*testrunner.NewTestRunner())

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.HandleFunc("/", pageHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/gh-org-checks.html")
	if err != nil {
		log.Error("Could not parse the html template!")
	}
	t.Execute(w, orgReport)
}

func initTestSuiteAndSchedule(testRunner testrunner.TestRunner) {
	testRunner.AddToTestSuites(testers.NewReleaseTester)
	testRunner.AddToTestSuites(testers.NewSecurityActionTester)
	testRunner.AddToTestSuites(testers.NewOSSTester)
	testRunner.AddToTestSuites(testers.NewTRGOneTester)
	testRunner.AddToTestSuites(testers.NewTRGTwoTester)

	scheduleCronJobs(testRunner)
}

func setLogLevel() {
	log.SetLevel(log.DebugLevel)
}

func scheduleCronJobs(testRunner testrunner.TestRunner) {
	log.Println("scheduled test cronjob")
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().Do(func() {
		go updateTestReport(testRunner)
	},
	)

	s.StartAsync()
}

func updateTestReport(testRunner testrunner.TestRunner) {
	log.Println("update test report")
	orgReport = testRunner.PerformRepoChecks()
}
