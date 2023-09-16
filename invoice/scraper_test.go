package invoice

import (
	"testing"
)

func FetchWebsite(t testing.TB) Invoice {
	invoice := Invoice{
		URL: "https://invoice.etax.nat.gov.tw",
	}

	success, err := invoice.Fetch()

	if !success {
		t.Error("Fetching website failed:", err)
	}

	return invoice
}
func TestFetchWebsite(t *testing.T) {
	FetchWebsite(t)
}

func TestGetJackpot(t *testing.T) {
	invoice := FetchWebsite(t)

	expectedCounts := 3
	prizes, err := invoice.GetJackpot()

	if err != nil {
		t.Error("parsing website failed:", err)
	}

	if len(prizes) != expectedCounts {
		t.Errorf("expect numbers of prize %d, got %d", expectedCounts, len(prizes))
	}
}
