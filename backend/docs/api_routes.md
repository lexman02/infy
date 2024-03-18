# API Documentation

## Overview
This document describes the API endpoints for the backend of our app.

## Endpoints
### Authentication Routes
The authentication routes handle user registration, login, and authentication status. For more details, refer to the `AuthRoutes` function in the `routes` package.

#### POST /auth/login
Authenticates a user and sets a JWT token in the response cookie.

**Request Example:**

```http
POST /auth/login
Content-Type: application/json

{
  "email": "test@test.com",
  "password": "testpass"
}
```

**Response Example:**

```json
{
  "success": "Logged in."
}
```

#### POST /auth/register
Registers a new user and sets a JWT token in the response cookie.

**Request Example:**

```http
POST /auth/register
Content-Type: application/json

{
  "email": "test@test.com",
  "password": "testpass",
  "confirm_password": "testpass",
  "username": "testuser",
  "first_name": "Test",
  "last_name": "User",
  "date_of_birth": "1990-01-01"
}
```

**Response Example:**

```json
{
  "success": "User created."
}
```

#### GET /auth/user
Fetches the currently authenticated user.

> ***Note:*** This route requires a valid JWT token in the request cookie.

**Request Example:**

```http
GET /auth/user
```

**Response Example:**

```json
{
  "user": {
    "username": "testuser",
    "email": "test@test.com",
    "isAdmin": false
  }
}
```


### Post Routes
The post routes handle creating, reading, updating, and deleting posts. For more details, refer to the `PostRoutes` function in the `routes` package.

#### GET /posts
Fetches all posts.

**Request Example:**

```http
GET /posts
```

**Response Example:**

```json
{
    [
        {
            "created": "2024-03-16 22:16:40",
            "post": {
                "id": "65f61a48db96c538c627ec5a",
                "user": {
                    "id": "65f2692e7533d263ef30e25a",
                    "username": "testuser",
                    "email": "test@test.com",
                    "isAdmin": false,
                    "profile": {
                        "first_name": "Test",
                        "last_name": "User",
                        "date_of_birth": "1990-01-01T00:00:00Z",
                        "avatar": "",
                        "rank": "Newbie",
                        "preferences": {
                            "genres": null,
                            "following": null,
                            "followers": null,
                            "watch_list": null,
                            "watched": null
                        }
                    }
                },
                "likes": 0,
                "dislikes": 0,
                "movie": {
                    "id": 1,
                    "title": "Movie Title 1",
                    "poster_path": "/path/to/poster1.jpg",
                    "tagline": "This is the tagline of the movie."
                },
                "content": "This is the content of the post."
            }
        },
        {
            "created": "2024-03-18 22:16:40",
            "post": {
                "id": "65f61a48db96c538c627ec5a",
                "user": {
                    "id": "65f2692e7533d263ef30e25a",
                    "username": "testuser",
                    "email": "test@test.com",
                    "isAdmin": false,
                    "profile": {
                        "first_name": "Test",
                        "last_name": "User",
                        "date_of_birth": "1990-01-01T00:00:00Z",
                        "avatar": "",
                        "rank": "Newbie",
                        "preferences": {
                            "genres": null,
                            "following": null,
                            "followers": null,
                            "watch_list": null,
                            "watched": null
                        }
                    }
                },
                "likes": 0,
                "dislikes": 0,
                "movie": {
                    "id": 2,
                    "title": "Movie Title 2",
                    "poster_path": "/path/to/poster2.jpg",
                    "tagline": "This is the tagline of the movie."
                },
                "content": "This is the content of the post."
            }
        }
    // More posts...
    ]
}
```

#### GET /posts/:id
Fetches a post by its ID.

**Parameters**
- `id` (string, required): The ID of the post to fetch.

**Request Example:**

```http
GET /posts/65f61a48db96c538c627ec5a
```

**Response Example:**

```json
{
    "created": "2024-03-16 22:16:40",
    "post": {
        "id": "65f61a48db96c538c627ec5a",
        "user": {
            "id": "65f2692e7533d263ef30e25a",
            "username": "testuser",
            "email": "test@test.com",
            "isAdmin": false,
            "profile": {
                "first_name": "Test",
                "last_name": "User",
                "date_of_birth": "1990-01-01T00:00:00Z",
                "avatar": "",
                "rank": "Newbie",
                "preferences": {
                    "genres": null,
                    "following": null,
                    "followers": null,
                    "watch_list": null,
                    "watched": null
                }
            }
        },
        "likes": 0,
        "dislikes": 0,
        "movie": {
            "id": 1,
            "title": "Movie Title 1",
            "poster_path": "/path/to/poster1.jpg",
            "tagline": "This is the tagline of the movie."
        },
        "content": "This is the content of the post."
    }
}
```

#### POST /posts
Creates a new post.

> ***Note:*** This route requires a valid JWT token in the request cookie.

**Request Example:**

```http
POST /posts
Content-Type: application/json

{
  "movie_id": 1,
  "content": "This is the content of the post."
}
```

**Response Example:**

```json
{
    "created": "2024-03-16 22:16:40",
    "post": {
        "id": "65f61a48db96c538c627ec5a",
        "user": {
            "id": "65f2692e7533d263ef30e25a",
            "username": "testuser",
            "email": "test@test.com",
            "isAdmin": false,
            "profile": {
                "first_name": "Test",
                "last_name": "User",
                "date_of_birth": "1990-01-01T00:00:00Z",
                "avatar": "",
                "rank": "Newbie",
                "preferences": {
                    "genres": null,
                    "following": null,
                    "followers": null,
                    "watch_list": null,
                    "watched": null
                }
            }
        },
        "likes": 0,
        "dislikes": 0,
        "movie": {
            "id": 1,
            "title": "Movie Title 1",
            "poster_path": "/path/to/poster1.jpg",
            "tagline": "This is the tagline of the movie."
        },
        "content": "This is the content of the post."
    }
}
```

#### PUT /posts/:id
Updates a post by its ID.

> ***Note:*** This route requires a valid JWT token in the request cookie.

**Parameters**
- `id` (string, required): The ID of the post to update.

**Request Example:**

```http
PUT /posts/65f61a48db96c538c627ec5a
Content-Type: application/json

{
  "content": "This is the updated content of the post."
}
```

**Response Example:**

```json
{
    "message": "Post updated successfully"
}
```

#### DELETE /posts/:id
Deletes a post by its ID.

> ***Note:*** This route requires a valid JWT token in the request cookie.

**Parameters**
- `id` (string, required): The ID of the post to delete.

**Request Example:**

```http
DELETE /posts/65f61a48db96c538c627ec5a
```

**Response Example:**

```json
{
    "message": "Post deleted successfully"
}
```


### Profile Routes
The profile routes handle fetching and updating user profiles. For more details, refer to the `ProfileRoutes` function in the `routes` package.

#### GET /profile/user
Fetches the currently authenticated user's profile.

> ***Note:*** This route requires a valid JWT token in the request cookie.

**Request Example:**

```http
GET /profile/user
```

**Response Example:**

```json
{
  "profile": {
    "first_name": "Test",
    "last_name": "User",
    "date_of_birth": "1990-01-01",
    "avatar": "",
    "rank": "Newbie",
    "preferences": {
      "genres": null,
      "following": null,
      "followers": null,
      "watch_list": null,
      "watched": null
    }
  }
}
```

#### TODO: PUT /profile/user
Updates the currently authenticated user's profile.

#### GET /profile/:username
Fetches a user's profile by their username.

**Parameters**
- `username` (string, required): The username of the user to fetch.

**Request Example:**

```http
GET /profile/testuser
```

**Response Example:**

```json
{
  "profile": {
    "first_name": "Test",
    "last_name": "User",
    "date_of_birth": "1990-01-01",
    "avatar": "",
    "rank": "Newbie",
    "preferences": {
      "genres": null,
      "following": null,
      "followers": null,
      "watch_list": null,
      "watched": null
    }
  }
}
```

#### POST /profile/follow/:id
Follows a user by their ID.

> ***Note:*** This route requires a valid JWT token in the request cookie.

**Parameters**
- `id` (string, required): The ID of the user to follow.

**Request Example:**

```http
POST /follow/65f2692e7533d263ef30e25a
```

**Response Example:**

```json
{
  "message": "User followed"
}
```

#### DELETE /profile/follow/:id
Unfollows a user by their ID.

> ***Note:*** This route requires a valid JWT token in the request cookie.

**Parameters**
- `id` (string, required): The ID of the user to unfollow.

**Request Example:**

```http
DELETE /unfollow/65f2692e7533d263ef30e25a
```

**Response Example:**

```json
{
  "message": "User unfollowed"
}
```

#### POST /profile/movies/add/watched
Adds a movie to the currently authenticated user's watched list.

> ***Note:*** This route requires a valid JWT token in the request cookie.

**Request Example:**

```http
POST /profile/movies/add/watched
Content-Type: application/json

{
  "movieId": 1
}
```

**Response Example:**

```json
{
  "message": "Movie added to watched list"
}
```

#### POST /profile/movies/add/watchlist
Adds a movie to the currently authenticated user's watch list.

> ***Note:*** This route requires a valid JWT token in the request cookie.

**Request Example:**

```http
POST /profile/movies/add/watchlist
Content-Type: application/json

{
  "movieId": 1
}
```

**Response Example:**

```json
{
  "message": "Movie added to watch list"
}
```

#### DELETE /profile/movies/remove/watched/:id
Removes a movie from the currently authenticated user's watched list.

> ***Note:*** This route requires a valid JWT token in the request cookie.

**Parameters**
- `id` (string, required): The ID of the movie to remove from the watched list.

**Request Example:**

```http
DELETE /profile/movies/remove/watched/1
```

**Response Example:**

```json
{
  "message": "Movie removed from watched list"
}
```

#### DELETE /profile/movies/remove/watchlist/:id
Removes a movie from the currently authenticated user's watch list.

> ***Note:*** This route requires a valid JWT token in the request cookie.

**Parameters**
- `id` (string, required): The ID of the movie to remove from the watch list.

**Request Example:**

```http
DELETE /profile/movies/remove/watchlist/1
```

**Response Example:**

```json
{
  "message": "Movie removed from watch list"
}
```

### Comment Routes
The comment routes handle creating, reading, updating, and deleting comments on posts. For more details, refer to the `CommentRoutes` function in the `routes` package.

> ***Note:*** Implementation of these routes is pending.

### TMDB Routes
The TMDB routes handle fetching movies from the external API.

#### GET /search/movies
Searches for movies in the external API.

**Request Example:**

```http
GET /search/movies?query=avengers
```

**Response Example:**

```json
{
  [
    {
      "id": 1,
      "title": "Movie Title 1",
      "poster_path": "/path/to/poster1.jpg",
      "release_date": "2022-01-01",
      "vote_average": 8.5,
      "overview": "This is a brief plot of Movie Title 1."
    },
    {
      "id": 2,
      "title": "Movie Title 2",
      "poster_path": "/path/to/poster2.jpg",
      "release_date": "2022-02-02",
      "vote_average": 7.5,
      "overview": "This is a brief plot of Movie Title 2."
    }
    // More movies...
  ]
}
```

#### GET /movies/:id
Fetches a movie's details by its ID from the external API.

** Parameters **
- `id` (string, required): The ID of the movie to fetch.

**Request Example:**

```http
GET /movies/1
```

**Response Example:**

```json
{
  "id": 1,
  "title": "Movie Title 1",
  "poster_path": "/path/to/poster.jpg",
  "tagline": "This is the tagline of the movie."
}
```