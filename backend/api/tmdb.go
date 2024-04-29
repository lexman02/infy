package api

import (
	"encoding/json"
	"fmt"
	"infy/models"
	"io"
	"net/http"
)

// Struct definitions for parsing JSON responses from TMDB API.
type TMDbMovieSearchResponse struct {
	Results []struct {
		ID         int     `json:"id"`
		Title      string  `json:"title"`
		PosterPath string  `json:"poster_path"`
		Year       string  `json:"release_date"`
		Rating     float64 `json:"vote_average"`
		Plot       string  `json:"overview"`
	} `json:"results"`
}

type TMDbPersonSearchResponse struct {
	Results []struct {
		ID          int    `json:"id"`
		ProfilePath string `json:"profile_path"`
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

// for actor page possibly
type TMDbActorDetailsResponse struct {
	Biography    string `json:"biography"`
	Birthday     string `json:"birthday"`
	Deathday     string `json:"deathday"`
	Gender       int    `json:"gender"`
	Name         string `json:"name"`
	PlaceOfBirth string `json:"place_of_birth"`
	ProfilePath  string `json:"profile_path"`
}

// For actor movies
type TMDbActorMovieCreditsResponse struct {
	Cast []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int     `json:"vote_count"`
		Popularity       float64 `json:"popularity"`
		Character        string  `json:"character"`
		CreditID         string  `json:"credit_id"`
	} `json:"cast"`
	ID int `json:"id"`
}

type TMDbVideoResponse struct {
	Results []struct {
		ID   string `json:"id"`
		Key  string `json:"key"`
		Name string `json:"name"`
		Site string `json:"site"` // YouTube or Vimeo
		Type string `json:"type"` // Trailer, Teaser, etc.
	} `json:"results"`
}

func SearchMovies(query string) (*TMDbMovieSearchResponse, error) {
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

	var searchResponse TMDbMovieSearchResponse
	if err := json.Unmarshal(body, &searchResponse); err != nil {
		return nil, err
	}

	return &searchResponse, nil
}

func SearchActors(query string) (*TMDbPersonSearchResponse, error) {
	apiKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/person?api_key=%s&query=%s", apiKey, query)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResponse TMDbPersonSearchResponse
	if err := json.Unmarshal(body, &searchResponse); err != nil {
		return nil, err
	}

	return &searchResponse, nil
}

// GetMovieDetails fetches detailed information about a specific movie from TMDB.
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

// IsValidMovieID checks if a given movie ID is valid by making an API call to TMDB.
func IsValidMovieID(movieID string) (bool, error) {
	tmdbAPIKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?api_key=%s", movieID, tmdbAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, nil
}

// GetTrendingMovies fetches trending movies from TMDB based on the specified time window (day or week).
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

// GetMovieCast fetches the cast list for a specific movie from TMDB.
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

// GetMovieReviews fetches reviews for a specific movie from TMDB.
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

// GetSimilarMovies fetches a list of movies similar to a specified movie from TMDB.
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

// for actor page
func GetActorDetails(actorID string) (*TMDbActorDetailsResponse, error) {
	apiKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/person/%s?api_key=%s", actorID, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var actorDetails TMDbActorDetailsResponse
	if err := json.Unmarshal(body, &actorDetails); err != nil {
		return nil, err
	}

	return &actorDetails, nil
}

// for actor details to show actos movies
func GetActorMovieCredits(actorID string) (*TMDbActorMovieCreditsResponse, error) {
	apiKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/person/%s/movie_credits?api_key=%s&language=en-US", actorID, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var movieCredits TMDbActorMovieCreditsResponse
	if err := json.Unmarshal(body, &movieCredits); err != nil {
		return nil, err
	}

	return &movieCredits, nil
}

func GetMovieTrailers(movieID string) (*TMDbVideoResponse, error) {
	apiKey := "89ab36f9a46f1199473c3da9950f2221"
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s/videos?api_key=%s&language=en-US", movieID, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var videoResponse TMDbVideoResponse
	if err := json.Unmarshal(body, &videoResponse); err != nil {
		return nil, err
	}

	return &videoResponse, nil
}
