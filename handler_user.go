package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/caesar-chin/go-rss-aggregator/internal/auth"
    "github.com/caesar-chin/go-rss-aggregator/internal/database"
    "github.com/google/uuid"
)

// handleCreateUser is a function to handle user creation requests
func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
    // Define the parameters expected in the request body
    type parameters struct {
        Name string `json:"name"`
    }
    // Create a new JSON decoder for the request body
    decoder := json.NewDecoder(r.Body)

    params := parameters{}
    // Decode the request body into the params struct
    err := decoder.Decode(&params)
    if err != nil {
        // If there's an error, respond with it
        respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
        return
    }

    // Create a new user in the database
    user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
        ID:        uuid.New(), // Generate a new UUID for the user
        CreatedAt: time.Now().UTC(), // Set the creation time to now
        UpdatedAt: time.Now().UTC(), // Set the update time to now
        Name:      params.Name, // Use the name from the request parameters
    })
    if err != nil {
        // If there's an error, respond with it
        respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %s", err))
        return
    }

    // Respond with the created user
    respondWithJSON(w, 200, databaseUserToUser(user))
}

// handleGetUserByAPIKey is a function to handle requests to get a user by their API key
func (apiCfg *apiConfig) handleGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
    // Extract the API key from the request header
    apiKey, err := auth.GetAPIKey(r.Header)

    if err != nil {
        // If there's an error, respond with it
        respondWithError(w, 403, fmt.Sprintf("Auth error: %s", err))
        return
    }

    // Get the user associated with the API key from the database
    user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

    if err != nil {
        // If there's an error, respond with it
        respondWithError(w, 403, fmt.Sprintf("Couldn't get user: %s", err))
        return
    }

    // Respond with the retrieved user
    respondWithJSON(w, 200, databaseUserToUser(user))
}