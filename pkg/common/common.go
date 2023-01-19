package common

import "github.com/catena-x/gh-org-checks/pkg/data"

func MergeReports(report1 data.RepositoryReport, report2 data.RepositoryReport) data.RepositoryReport {
	mergedReport := report1

	mergedReport.TestSucceed = mergedReport.TestSucceed && mergedReport.TestSucceed
	mergedReport.Log = append(mergedReport.Log, report2.Log...)

	return mergedReport
}
