package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/thejasms1603/rssagg/internal/database"
)


func (apiCfg *apiConfig) hanldeFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}


	decode := json.NewDecoder(r.Body)


	params := parameters{}
	err := decode.Decode(&params)
	if err != nil {
		respondWithError(w,400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}


	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID: params.FeedID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w,400, fmt.Sprintf("Error creating feed follow: %s", err))
		return
	}
	respondWithJSON(w,201, databaseFeedFollowToFeedFollow(feedFollow))
}


func (apiCfg *apiConfig) handleGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w,400, fmt.Sprintf("Error getting feed follows: %s", err))
		return
	}
	respondWithJSON(w,200, databaseFeedFollowsToFeedFollows(feedFollows))
}



func (apiCfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowId := chi.URLParam(r, "feedFollowId")
	feedFollowIdUUID, err := uuid.Parse(feedFollowId)
	if err != nil {
		respondWithError(w,400, fmt.Sprintf("Error parsing feed follow id: %s", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowIdUUID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w,400, fmt.Sprintf("Error deleting feed follow: %s", err))
		return
	}
	respondWithJSON(w,200, struct{}{})
}