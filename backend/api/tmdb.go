package api

import (
	"encoding/json"
	"fmt"
	"infy/models"
	"io"
	"net/http"
)

type TMDbSearchResponse struct {
	Results []struct {
		ID         int     `json:"id"`
		Title      string  `json:"title"`
		PosterPath string  `json:"poster_path"`
		Year       string  `json:"release_date"`
		Rating     float64 `json:"vote_average"`
		Plot       string  `json:"overview"`
	} `json:"results"`
}

func SearchMovies(query string) (*TMDbSearchResponse, error) {
	apiKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s", apiKey, query)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResponse TMDbSearchResponse
	if err := json.Unmarshal(body, &searchResponse); err != nil {
		return nil, err
	}

	return &searchResponse, nil
}

func GetMovieDetails(movieID string) (*models.Movie, error) {
	apiKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?api_key=%s", movieID, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var movieDetails models.Movie
	if err := json.Unmarshal(body, &movieDetails); err != nil {
		return nil, err
	}

	return &movieDetails, nil
}

func IsValidMovieID(movieID string) (bool, error) {
	tmdbAPIKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?api_key=%s", movieID, tmdbAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return false, err // Network error or issue with request
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil // Movie ID is valid
	}
	return false, nil // Movie ID is not valid, but no error occurred
}
