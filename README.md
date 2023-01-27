# gh-org-checks

## What is the service

The service will check on each and every repository daily (see [ticket](https://jira.catena-x.net/browse/A1ODT-504)), the available checks are:


- **releaseCheck**:
  check if either of two condition are met
    - if a `changelog.md` file exist
    - if a GitHub release exist and if so, weather it use semantic versioning


- **OSSCheck**: Check if all necessary files are present for repository to be open source ready. Currently, this is 
  applicable to repositories under organization "Catena-X", see documentation [here](https://github.com/catenax-ng/foss-example).
- **SecurityActionCheck**: Check if list of GitHub action suggested by security team are present
- **TRG 1 Check**: Check of the repository is compliance with eclipse release guideline (o.e. TRG) One
  - [TRG 1.01 : README.md](https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-1)
  - [TRG 1.03: CHANGELOG.md](https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-3)
- **TRG 2 Check**: Check of the repository is compliance with TRG Two
  - [TRG 2.01 : Default Branch](https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-1)
  - [TRG 2.03: Repo structure](https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-3)

## How to run the service
The service will run checks across GitHub repositories from a given GitHub organization

set the following environment variables:

- GITHUB_ACCESS_TOKEN
  
Contain the github access token, which has admin permission to the targeted GitHub organization

- GITHUB_ORG_NAME 

Contain the name of github org, if unset, default to "catenax-ng"


## How to provide more checks

You can always extend and add new tests, please be aware that you should

- Implement interface [GithubTester](pkg/testers/githubTester.go) and [testProperty](pkg/testers/testProperty.go), implement the function `PerformTest` and `GetTestName`, the functions will be called from [testrunner](pkg/testrunner/testRunner.go)
- If you want to check a file/folder exist in the repository, you can reuse [contentTester](pkg/testers/contentTester.go) as the base struct, use property `contents` to indicate the file/folder you want to check
- Add your test to testsuite (i.e. to be run) at [main](main.go) on function `initTestSuiteAndSchedule`
