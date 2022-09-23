# gh-org-checks

## What is the service

The service will check on each and every repository daily (see [ticket](https://jira.catena-x.net/browse/A1ODT-504)), the available checks are:

- readmeCheck:

  check if a readme file is present 

- helmChartCheck:

  check if the directory `chart` is present under root directory

- releaseCheck:
  check if either of two condition are met
    - if a `changelog.md` file exist
    - if a github release exist and if so, weather it use semantic versioning


- OSSCheck:
  Check if all necessary files are present for repository to be open source ready

- SecurityActionCheck:
- Check if list of github action suggested by security team are present


## How to run the service
The service will run checks across Github repositories from a given Github organization

set the following environment variables:

- GITHUB_ACCESS_TOKEN
  
Contain the github access token, which has admin permission to the targeted Github organization

- GITHUB_ORG_NAME 

Contain the name of github org, if unset, default to "catenax-ng"


## How to provide more checks

You can always extend and add new tests, please be aware that you should

- Implement interface [GithubTester](pkg/testers/githubTester.go), implement the function `PerformTest`, the function will be called from [testrunner](pkg/testrunner/testRunner.go)
- If you want to check a file/folder exist in the repository, you can reuse [contentTester](pkg/testers/contentTester.go), use property `contents` to indicate the file/folder you want to check
- Add your test to testsuite (i.e. to be run) at [main](main.go) on function `initTestSuiteAndSchedule`
