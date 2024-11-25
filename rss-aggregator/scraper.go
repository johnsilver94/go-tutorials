package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/wagslane/rssagg/internal/database"
)

func startScraping(db *database.Queries, concurrency int, waitTime time.Duration) {
	log.Printf("Starting scraping with %v workers every %s ", concurrency, waitTime)

	ticker := time.NewTicker(waitTime)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error fetching feeds: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}

		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		log.Printf("Found item: %v on feed %v", item.Title, feed.Name)
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing time: %v", err)
			continue

		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Description: description,
			Url:         item.Link,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		})

		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				continue
			}
			log.Printf("Error creating post: %v", err)
		}
	}

	log.Printf("Scraped feed: %v , %v posts found", feed.Name, len(rssFeed.Channel.Items))
}
