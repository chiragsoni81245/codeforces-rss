package feed

import (
	"fmt"
	"github.com/gorilla/feeds"
	"time"
)

// GenerateRSS generates an RSS feed with a random Codeforces question
func GenerateRSS(tags []string, excludedTags []string, minRating, maxRating int) (string, error) {
	// Fetch a problem
	problem, err := FetchRandomProblem(tags, excludedTags, minRating, maxRating)
	if err != nil {
		return "", err
	}

	// Create an RSS feed
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Daily Codeforces Problem",
		Link:        &feeds.Link{Href: "https://codeforces.com"},
		Description: "A randomly selected problem from Codeforces based on your preferred tags.",
		Created:     now,
	}

	// Add problem as an RSS item
	feed.Items = []*feeds.Item{
		{
			Title:       fmt.Sprintf("%s (%d)", problem.Name, problem.Rating),
			Link:        &feeds.Link{Href: fmt.Sprintf("https://codeforces.com/problemset/problem/%d/%s", problem.ContestID, problem.Index)},
			Description: fmt.Sprintf("Tags: %v", problem.Tags),
			Created:     now,
		},
	}

	// Convert to RSS format
	return feed.ToRss()
}

