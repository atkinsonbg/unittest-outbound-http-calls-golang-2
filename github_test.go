package github

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitHubCallSuccess(t *testing.T) {

	// build our response JSON
	jsonResponse := `[{
			"full_name": "mock-repo"
		}]`

	// create a new reader with that JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	ghm := GitHubManager{
		BaseUrl: server.URL,
	}

	result, err := ghm.GetRepos("atkinsonbg")
	if err != nil {
		t.Error("TestGitHubCallSuccess failed.")
		return
	}

	if len(result) == 0 {
		t.Error("TestGitHubCallSuccess failed, array was empty.")
		return
	}

	if result[0]["full_name"] != "mock-repo" {
		t.Error("TestGitHubCallSuccess failed, array was not sorted correctly.")
		return
	}
}

func TestGitHubCallFail(t *testing.T) {
	// create a new reader with that JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}))
	defer server.Close()

	ghm := GitHubManager{
		BaseUrl: server.URL,
	}

	_, err := ghm.GetRepos("atkinsonbgthisusershouldnotexist")
	if err == nil {
		t.Error("TestGitHubCallFail failed.")
		return
	}

	assert.Equal(t, "EOF", err.Error())
}
