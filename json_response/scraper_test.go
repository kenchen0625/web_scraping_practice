package jsonresponse

import "testing"

const (
	GetUsersUrl   = "https://jsonplaceholder.typicode.com/users"
	GetUserOneUrl = "https://jsonplaceholder.typicode.com/users/1"
)

func FetchWebsite(t testing.TB, url string) UsersFetcher {
	usersFetcher := UsersFetcher{
		URL: url,
	}
	err := usersFetcher.Fetch()

	if err != nil {
		t.Error("GET users/ error:", err)
	}

	return usersFetcher
}

func TestFetch(t *testing.T) {
	FetchWebsite(t, GetUsersUrl)
}

func TestParseUsers(t *testing.T) {
	usersFetcher := FetchWebsite(t, GetUsersUrl)
	_, err := usersFetcher.GetUsers()
	if err != nil {
		t.Error(err)
	}
}

func TestParseOneUser(t *testing.T) {
	usersFetcher := FetchWebsite(t, GetUserOneUrl)
	_, err := usersFetcher.GetOneUser()
	if err != nil {
		t.Error(err)
	}
}
