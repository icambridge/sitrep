package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/icambridge/gobucket"
	"sitrep/app/services"
	"sitrep/app/converters"
	"sitrep/app/entities"
	"strings"
)

type PullRequests struct {
	App
}

func (c PullRequests) All() revel.Result {

	pullRequests := []*gobucket.PullRequest{}
	bitbucket := services.GetBitbucket()
	bitbucketOwner := revel.Config.StringDefault("bitbucket.owner", "")

	repos, err := bitbucket.Repositories.GetAll(bitbucketOwner)

	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		parts := strings.Split(repo.FullName, "/")
		newRequests, err := bitbucket.PullRequests.GetAll(bitbucketOwner, parts[1])
		if err != nil {
			panic(err)
		}
		pullRequests = append(pullRequests, newRequests...)
	}
	convertedPrs := converters.GobucketToSitRepMultiPrs(pullRequests)

	entity := entities.RepoInfo{PullRequests: convertedPrs}

	jsonData, _ := json.Marshal(entity)
	json := string(jsonData)
	c.Request.Format = "json"
	return c.Render(json)
}
