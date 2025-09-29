package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/thejasms1603/rssagg/internal/database"
)

type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name string `json:"name"`
	APIKey string `json:"ApiKey"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		Name: dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		APIKey: dbUser.ApiKey,
	}
}


type Feed struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Url string `json:"url"`
	UserID uuid.UUID `json:"userId"`
}


func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		Name: dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func databaseFeedsTOFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}


type FeedFollow struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FeedID uuid.UUID `json:"feedId"`
	UserID uuid.UUID `json:"userId"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		FeedID: dbFeedFollow.FeedID,
		UserID: dbFeedFollow.UserID,
	}
}


func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, feedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(feedFollow))
	}
	return feedFollows

}