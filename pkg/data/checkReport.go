package data

type TestStatus string

const (
	Successful TestStatus = "Successful"
	Failed     TestStatus = "Failure"
)

type OrgReports struct {
	OrgName             string                `json:"GithubOrgName,omitempty"`
	Error               string                `json:"ErrorMessage,omitempty"`
	RepositoriesReports []RepositoriesReports `json:"RepositoriesReports,omitempty"`
}

type RepositoriesReports struct {
	RepositoryName   string             `json:"RepositoryName,omitempty"`
	RepositoryReport []RepositoryReport `json:"RepositoryReport,omitempty"`
}

type RepositoryReport struct {
	TestName    string     `json:"TestName,omitempty"`
	GithubRepo  string     `json:"GithubRepositoryName,omitempty"`
	CheckStatus TestStatus `json:"TestStatus,omitempty"`
	Log         string     `json:"Log,omitempty"`
}
