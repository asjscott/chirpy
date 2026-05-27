package main

import "net/http"

func (cfg *apiConfig) handlerListChirps(w http.ResponseWriter, r * http.Request) {
	chirpsResponse := []Chirp{}

	chirps, err := cfg.db.ListChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	for _, chirp := range chirps {
		val := Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
		}
		chirpsResponse = append(chirpsResponse, val)
	}

	respondWithJSON(w, http.StatusOK, chirpsResponse)

}