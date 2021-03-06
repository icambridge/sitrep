package entities


import (
	"sitrep/app/models"
)

type PullRequestList struct {
	Values []PullRequest `json:"values"`
}

type PullRequest struct {
	Title       string       `json:"title"`
	Author      string       `json:"author"`
	Source      string       `json:"source"`
	Destination string       `json:"destination"`
	Id          int          `json:"id"`
	Approvals   []Approval   `json:"approvals"`
	Url         string       `json:"url"`
	LastBuild   models.Build `json:"last_build"`
}

type Approval struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

type Branch struct {
	Name    string           `json:"name"`
	BuildId int              `json:"build_id"`
	Status  string           `json:"status"`
	Timestamp int64          `json:"timestamp"`
	LastBuild   models.Build `json:"last_build"`
}

type RepoInfo struct {
	Name         string        `json:"name"`
	PullRequests []PullRequest `json:"pull_requests"`
	Branches     []Branch      `json:"branches"`
}
