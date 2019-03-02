package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TagsResponse is Response structure for Github Tags API
type TagsResponse []Tag

// Tag is ...
type Tag struct {
	Name       string `json:"name"`
	Commit     Commit `json:"commit"`
	ZipballURL string `json:"zipball_url"`
	TarballURL string `json:"tarball_url"`
}

// Commit is ...
type Commit struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

// Fetch is ...
func Fetch(owner, repo string) *TagsResponse {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/tags", owner, repo)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var result TagsResponse
	if err := json.Unmarshal(b, &result); err != nil {
		panic(err)
	}
	return &result
}

func (data *TagsResponse) GetLatestTag(checker VersionChecker) *Tag {
	tag := (*data)[0]
	rest := (*data)[1:]

	for _, d := range rest {
		if checker.Check(tag.Name, d.Name) {
			tag = d
		}
	}
	return &tag
}

// VersionChecker is ...
type VersionChecker interface {
	Check(nowVersion, newVersion string) bool
}
