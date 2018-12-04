// package quickumlsrest provides an API client to the quickUMLS-rest service.
// The server can be found as a docker image from https://hub.docker.com/r/aehrc/quickumls-rest/.
package quickumlsrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

// Candidate defines the response from the quickUMLS-rest service.
type Candidate struct {
	Start      int      `json:"start"`
	End        int      `json:"end"`
	Preferred  int      `json:"preferred"`
	Similarity float64  `json:"similarity"`
	CUI        string   `json:"cui"`
	NGram      string   `json:"ngram"`
	Term       string   `json:"term"`
	SemTypes   []string `json:"semtypes"`
	SnomedCT   []string `json:"snomed_ct"`
}

// Candidates is a collection of candidates.
type Candidates []Candidate

// MatchRequest is the request body sent to the match endpoint of the quickUMLS-rest service.
type MatchRequest struct {
	Text string `json:"text"`
}

// Client is a client to the quickUMLS-rest service.
type Client struct {
	URL string
	*http.Client
}

// NewClient creates a new client for the service.
func NewClient(URL string) Client {
	return Client{
		URL:    URL,
		Client: &http.Client{},
	}
}

// Match performs a match request and returns a slice of candidates.
func (c Client) Match(text string) (Candidates, error) {
	mr := MatchRequest{
		Text: text,
	}
	b, err := json.Marshal(mr)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(c.URL)
	u.Path = path.Join(u.Path, "match")
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var candidates Candidates
	err = json.NewDecoder(resp.Body).Decode(&candidates)
	if err != nil {
		return nil, err
	}

	return candidates, nil
}
