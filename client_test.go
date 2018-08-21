package quickumlsrest_test

import (
	"testing"
	"github.com/hscells/quickumlsrest"
)

func TestMatch(t *testing.T) {
	client := quickumlsrest.NewClient("http://localhost:5000")
	candidates, err := client.Match("cancer")
	if err != nil {
		t.Fatal(err)
	}

	for _, candidate := range candidates {
		t.Log(candidate.CUI)
	}
}
