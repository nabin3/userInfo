package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/nabin3/userInfo/proto"
)

const (
	Min_User_Height = 3
	Max_User_Height = 7
)

// Handler for "POST /adduser"
func (cfg *apiConfig) handlerAddUser(w http.ResponseWriter, r *http.Request) {
	// Type for recieving data from request body
	type parameters struct {
		Fname     string  `json:"fname"`
		City      string  `json:"city"`
		Phone     string  `json:"phone"`
		Height    float32 `json:"height"`
		IsMarried bool    `json:"is_married"`
	}

	// Decoding recieved json data
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("%v", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't decode recieved user-data")
		return
	}

	// Data validation
	if params.Height < Min_User_Height && params.Height > Max_User_Height {
		respondWithError(w, http.StatusBadRequest, "give reasonable height")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Adding a user
	user_id, err := cfg.client.AddUser(ctx, &pb.User{
		Fname:     params.Fname,
		City:      params.City,
		Phone:     params.Phone,
		Height:    params.Height,
		Ismarried: params.IsMarried,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't add user")
		return
	}

	respondWithJson(w, http.StatusOK, generatedUserIdToResponseUserId(user_id))
}

// Handler for "GET /getuser"
func (cfg *apiConfig) handlerRetrieveSingleUser(w http.ResponseWriter, r *http.Request) {
	// Format of recieved data from request body
	type parameters struct {
		UserId string `json:"user_id"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode recieved user_id")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Retrieveing one user
	user, err := cfg.client.RetrieveOneUser(ctx, &pb.UserID{Id: params.UserId})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "couldn't find the user")
		return
	}

	respondWithJson(w, http.StatusOK, retrievedUserToResponseUser(user))
}

// Handler for "GET /get_multiple_users"
func (cfg *apiConfig) handlerRetrieveMultipleUsers(w http.ResponseWriter, r *http.Request) {
	// Format of recieved data from request body
	type parameters struct {
		UserIdList []string `json:"user_id_list"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("error in handlerRetrieveMultipleUser func at decoding recieved data: %v", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't decode recieved user_id_list")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Retrieving users
	usersList, err := cfg.client.RetrieveMultipleUsers(ctx, &pb.UserIDList{Ids: params.UserIdList})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}

	respondWithJson(w, http.StatusOK, retrievedUserListToResponseUserList(usersList))
}

// Handler for "GET /search_users"
func (cfg *apiConfig) handlerSearchUsers(w http.ResponseWriter, r *http.Request) {
	// Format of recieved data from request body
	type parameters struct {
		SearchCriteriaName string      `json:"search_criteria"`
		Value              interface{} `json:"value"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not decode recieved data")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Find user base on search_criteria
	switch params.SearchCriteriaName {
	case "fname":
		userList, err := cfg.client.SearchUsers(ctx, &pb.UserSearchCriteria{Fname: fmt.Sprintf("%v", params.Value)})
		if err != nil {
			respondWithError(w, http.StatusContinue, "didn't find any users")
			return
		}
		respondWithJson(w, http.StatusOK, retrievedUserListToResponseUserList(userList))

	case "city":
		userList, err := cfg.client.SearchUsers(ctx, &pb.UserSearchCriteria{City: fmt.Sprintf("%v", params.Value)})
		if err != nil {
			respondWithError(w, http.StatusContinue, "didn't find any users")
			return
		}
		respondWithJson(w, http.StatusOK, retrievedUserListToResponseUserList(userList))

	case "phone":
		userList, err := cfg.client.SearchUsers(ctx, &pb.UserSearchCriteria{Phone: fmt.Sprintf("%v", params.Value)})
		if err != nil {
			respondWithError(w, http.StatusContinue, "didn't find any users")
			return
		}
		respondWithJson(w, http.StatusOK, retrievedUserListToResponseUserList(userList))

	case "height":
		height, ok := params.Value.(float32)
		if !ok {
			respondWithError(w, http.StatusBadRequest, "pass valid height")
			return
		}
		userList, err := cfg.client.SearchUsers(ctx, &pb.UserSearchCriteria{Height: height})
		if err != nil {
			respondWithError(w, http.StatusContinue, "didn't find any users")
			return
		}
		respondWithJson(w, http.StatusOK, retrievedUserListToResponseUserList(userList))

	case "is_married":
		is_married, ok := params.Value.(bool)
		if !ok {
			respondWithError(w, http.StatusBadRequest, "if want to find users who are married pass true as search phrase")
			return
		}
		userList, err := cfg.client.SearchUsers(ctx, &pb.UserSearchCriteria{Ismarried: is_married})
		if err != nil {
			respondWithError(w, http.StatusContinue, "didn't find any users")
			return
		}
		respondWithJson(w, http.StatusOK, retrievedUserListToResponseUserList(userList))

	default:
		respondWithError(w, http.StatusBadRequest, "please give proper search_criteria")
	}
}
