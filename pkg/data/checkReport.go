package data

type OrgReports struct {
	OrgName           string                `json:"GithubOrgName,omitempty"`
	LastTestTime      string                `json:"LastTestTime,omitempty"`
	NumOfRepos        int                   `json:"NumOfRepos,omitempty"`
	Error             error                 `json:"ErrorMessage,omitempty"`
	RepositoryReports []RepositoriesReports `json:"RepositoryReports,omitempty"`
}

type RepositoriesReports struct {
	RepositoryName   string             `json:"RepositoryName,omitempty"`
	RepositoryURL    string             `json:"RepositoryURL,omitempty"`
	RepositoryReport []RepositoryReport `json:"RepositoryReport,omitempty"`
}

type RepositoryReport struct {
	TestName    string   `json:"TestName,omitempty"`
	GithubRepo  string   `json:"GithubRepositoryName,omitempty"`
	TestSucceed bool     `json:"TestSucceed,omitempty"`
	Log         []string `json:"Log,omitempty"`
}
