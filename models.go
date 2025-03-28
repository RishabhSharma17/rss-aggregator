package main

import (
	"time"

	"github.com/RishabhSharma17/rssaggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey	  string 	`json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url 	  string 	`json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func DatabaseUserToUser(user database.User) User{
	return User{
		ID: 	  user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:     user.Name,
		APIKey:   user.ApiKey,
	}
}

func DatabaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url: 	  feed.Url,
		UserID:    feed.UserID,
	}
}

func DatabaseFeedsToFeeds(feed []database.Feed) []Feed {
	feeds := []Feed{}
	for _,f := range feed{
		feeds = append(feeds, DatabaseFeedToFeed(f))
	}
	return feeds
}

func DatabaseFeedFollowToFeedFollow(feedFollow database.FeedFollow) FeedFollow{
	return FeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		FeedID:    feedFollow.FeedID,
		UserID:    feedFollow.UserID,
	}
}

func DatabaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow{
	feedFollows := []FeedFollow{}
    for _,dbFeedFollow := range dbFeedFollows{
		feedFollows = append(feedFollows,DatabaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}