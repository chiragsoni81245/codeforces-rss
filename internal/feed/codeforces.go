package feed

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

const codeforcesAPI = "https://codeforces.com/api/problemset.problems"

// Problem represents a Codeforces problem
type Problem struct {
	ContestID int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Tags      []string `json:"tags"`
	Rating    int      `json:"rating"`
}

// APIResponse represents the Codeforces API response
type APIResponse struct {
	Result struct {
		Problems []Problem `json:"problems"`
	} `json:"result"`
}

// FetchRandomProblem fetches a random problem based on difficulty & tags
func FetchRandomProblem(tags []string, minRating, maxRating int) (*Problem, error) {
	resp, err := http.Get(codeforcesAPI)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch problems: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data APIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// Filter problems by tags & difficulty
	var filtered []Problem
	for _, problem := range data.Result.Problems {
		if problem.Rating >= minRating && problem.Rating <= maxRating {
			for _, tag := range tags {
				for _, pTag := range problem.Tags {
					if tag == pTag {
						filtered = append(filtered, problem)
						break
					}
				}
			}
		}
	}

	if len(filtered) == 0 {
		return nil, fmt.Errorf("no problems found for the given filters")
	}

	// Pick a random problem
    randomNumberGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &filtered[randomNumberGenerator.Intn(len(filtered))], nil
}

