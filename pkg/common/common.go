package common

import "github.com/catena-x/gh-org-checks/pkg/data"

const TestTrgOne = "TRG 1"
const TestTrgTwo = "TRG 2"

const TrgOneOne = "TRG 1.1"
const TrgOneTwo = "TRG 1.2"
const TrgOneThree = "TRG 1.3"
const TrgTwoOne = "TRG 2.1"
const TrgTwoTwo = "TRG 2.2"
const TrgTwoThree = "TRG 2.3"

var refMap = make(map[string]string)

func init() {
	refMap[TrgOneOne] = "https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-1"
	refMap[TrgOneTwo] = "https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-2"
	refMap[TrgOneThree] = "https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-3"
	refMap[TrgTwoOne] = "https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-1"
	refMap[TrgTwoTwo] = "https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-2"
	refMap[TrgTwoThree] = "https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-3"
}

func GetRefLink(key string) string {
	return refMap[key]
}
func MergeReports(firstReport data.RepositoryReport, moreReports ...data.RepositoryReport) data.RepositoryReport {
	mergedReport := firstReport

	for _, report := range moreReports {
		mergedReport.TestSucceed = mergedReport.TestSucceed && report.TestSucceed
		mergedReport.Log = append(mergedReport.Log, report.Log...)
	}

	return mergedReport
}
