package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Label struct {
	Id          int    `json:"id"`
	NodeId      string `json:"node_id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Def         bool   `json:"default"`
	Description string `json:"description"`
}

var client = http.Client{}

func getRepositoryLabels(owner, repository, token string) []Label {
	req := http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: "https",
			Host:   "api.github.com",
			Path:   fmt.Sprintf("/repos/%v/%v/labels", owner, repository),
		},
		Header: http.Header{"Authorization": []string{"Bearer " + token}},
	}

	res, _ := client.Do(&req)
	body, _ := ioutil.ReadAll(res.Body)

	var labels []Label
	_ = json.Unmarshal(body, &labels)

	return labels
}

func deleteRepositoryLabels(owner, repository, token string) {
	labels := getRepositoryLabels(owner, repository, token)

	for _, label := range labels {
		req := http.Request{
			Method: http.MethodDelete,
			URL: &url.URL{
				Scheme: "https",
				Host:   "api.github.com",
				Path:   fmt.Sprintf("/repos/%v/%v/labels/%v", owner, repository, label.Name),
			},
			Header: http.Header{"Authorization": []string{"Bearer " + token}},
		}

		res, _ := client.Do(&req)
		if res.StatusCode == http.StatusNoContent {
			println("Deleted label: " + label.Name)
		}
	}
}

func writeRepositoryLabels(owner, repository, token string, labels []Label) {
	for _, label := range labels {
		json_label, _ := json.Marshal(label)

		req := http.Request{
			Method: http.MethodPost,
			URL: &url.URL{
				Scheme: "https",
				Host:   "api.github.com",
				Path:   fmt.Sprintf("/repos/%v/%v/labels", owner, repository),
			},
			Header: http.Header{"Authorization": []string{"Bearer " + token}},
			Body:   io.NopCloser(bytes.NewReader(json_label)),
		}

		res, _ := client.Do(&req)
		if res.StatusCode == http.StatusCreated {
			println("Created label: " + label.Name)
		}
	}
}

func readCommandLineArguments() (string, string, string, string) {
	if len(os.Args) != 5 {
		log.Fatal("Usage: go run <owner> <repository> <token> <labels>")
	}

	source_owner := os.Args[1]
	source_repo := os.Args[2]
	target_owner := os.Args[3]
	target_repo := os.Args[4]

	return source_owner, source_repo, target_owner, target_repo
}

func main() {
	source_owner, source_repo, target_owner, target_repo := readCommandLineArguments()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	gh_token := os.Getenv("GH_TOKEN")

	source_labels := getRepositoryLabels(source_owner, source_repo, gh_token)
	deleteRepositoryLabels(target_owner, target_repo, gh_token)
	writeRepositoryLabels(target_owner, target_repo, gh_token, source_labels)
}
