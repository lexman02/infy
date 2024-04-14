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

type TMDbTrendingResponse struct {
	Results []struct {
		ID         int     `json:"id"`
		Title      string  `json:"title"`
		PosterPath string  `json:"poster_path"`
		Year       string  `json:"release_date"`
		Rating     float64 `json:"vote_average"`
		Plot       string  `json:"overview"`
	} `json:"results"`
}

type TMDbCastResponse struct {
	Cast []struct {
		CastID      int    `json:"cast_id"`
		Character   string `json:"character"`
		Name        string `json:"name"`
		ProfilePath string `json:"profile_path"`
	} `json:"cast"`
}

type TMDbReviewResponse struct {
	Results []struct {
		Author  string `json:"author"`
		Content string `json:"content"`
		ID      string `json:"id"`
		URL     string `json:"url"`
	} `json:"results"`
}

type TMDbSimilarMoviesResponse struct {
	Results []struct {
		ID         int    `json:"id"`
		Title      string `json:"title"`
		PosterPath string `json:"poster_path"`
		Overview   string `json:"overview"`
	} `json:"results"`
}

type MovieDetails struct {
	models.Movie
	Overview     string `json:"overview"`
	BackdropPath string `json:"backdrop_path"`
	Runtime      int    `json:"runtime"`
	ReleaseDate  string `json:"release_date"`
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

func GetMovieDetails(movieID string, limited bool) (*MovieDetails, *models.Movie, error) {
	apiKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?api_key=%s", movieID, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var movieDetails MovieDetails
	if err := json.Unmarshal(body, &movieDetails); err != nil {
		return nil, nil, err
	}

	if limited {
		movie := models.Movie{
			ID:         movieDetails.ID,
			Title:      movieDetails.Title,
			PosterPath: movieDetails.PosterPath,
			Tagline:    movieDetails.Tagline,
		}

		return nil, &movie, nil
	}

	return &movieDetails, nil, nil
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

// timeWindow can be "day" or "week".
func GetTrendingMovies(timeWindow string) (*TMDbTrendingResponse, error) {
	apiKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/trending/movie/%s?api_key=%s", timeWindow, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var trendingResponse TMDbTrendingResponse
	if err := json.Unmarshal(body, &trendingResponse); err != nil {
		return nil, err
	}

	return &trendingResponse, nil
}

// Fetches cast information for a movie
func GetMovieCast(movieID string) (*TMDbCastResponse, error) {
	apiKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s/credits?api_key=%s", movieID, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var castResponse TMDbCastResponse
	if err := json.Unmarshal(body, &castResponse); err != nil {
		return nil, err
	}

	return &castResponse, nil
}

// Fetches reviews for a movie
func GetMovieReviews(movieID string) (*TMDbReviewResponse, error) {
	apiKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s/reviews?api_key=%s", movieID, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var reviewsResponse TMDbReviewResponse
	if err := json.Unmarshal(body, &reviewsResponse); err != nil {
		return nil, err
	}

	return &reviewsResponse, nil
}

// Fetches similar movies
func GetSimilarMovies(movieID string) (*TMDbSimilarMoviesResponse, error) {
	apiKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s/similar?api_key=%s", movieID, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var similarMoviesResponse TMDbSimilarMoviesResponse
	if err := json.Unmarshal(body, &similarMoviesResponse); err != nil {
		return nil, err
	}

	return &similarMoviesResponse, nil
}
