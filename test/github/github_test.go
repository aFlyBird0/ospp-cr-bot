package github

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/cbrgm/githubevents/githubevents"
	"github.com/google/go-github/v44/github"
)

func TestGithub(t *testing.T) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		token = os.Getenv("github_token")
	}
	if token == "" {
		t.Error("GITHUB_TOKEN not set")
	}
	client := github.NewClient(nil)

	// list all organizations for union "willnorris"
	orgs, _, err := client.Organizations.List(context.Background(), "aFlyBird0", nil)
	if err != nil {
		t.Error(err)
	}
	for _, org := range orgs {
		t.Logf("%s-%s\n", org.GetLogin(), org.GetURL())
	}

	t.Log("\n")

	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "hduhelp", opt)
	if err != nil {
		t.Error(err)
	}
	for _, repo := range repos {
		t.Logf("%s-%s\n", repo.GetName(), repo.GetURL())
	}
}

func TestGithubWebhook(t *testing.T) {

	// create a new event handler
	handle := githubevents.New("123456")

	// add callbacks
	handle.OnIssueCommentCreated(
		func(deliveryID string, eventName string, event *github.IssueCommentEvent) error {
			fmt.Printf("%#v made a comment!\n", event.Sender)
			fmt.Printf("%#v\n", event.Comment)
			return nil
		},
	)

	handle.OnPullRequestEventAny(
		func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
			fmt.Printf("%s made a pull request!", event.Sender.Login)
			return nil
		},
	)

	// add a http handleFunc
	http.HandleFunc("/hook", func(w http.ResponseWriter, r *http.Request) {
		err := handle.HandleEventRequest(r)
		if err != nil {
			fmt.Println("error")
		}
	})

	// start the server listening on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
