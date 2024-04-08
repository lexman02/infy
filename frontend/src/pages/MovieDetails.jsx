import React, { useEffect, useState, useContext } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import { UserContext } from '../contexts/UserProvider';

export default function MovieDetails() {
  const { movieID } = useParams();
  const [movie, setMovie] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const { userData } = useContext(UserContext); 

  useEffect(() => {
    if (movieID) {
      setIsLoading(true);
      axios.get(`http://localhost:8000/movies/${movieID}`)
        .then(response => {
          setMovie(response.data);
          setIsLoading(false);
        })
        .catch(error => {
          console.error("Failed to fetch movie details:", error);
          setIsLoading(false);
        });
    } else {
      console.log("Movie ID is undefined.");
    }
  }, [movieID]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (!movie) {
    return <div>Movie not found.</div>;
  }

  return (
    <div className="text-neutral-100 font-sans p-4">
      <div className="flex flex-col items-center">
        <img src={`https://image.tmdb.org/t/p/original/${movie.poster_path}`} alt={movie.title} className="w-60 h-90 object-cover rounded-lg mb-4"/>
        <h1 className="text-3xl font-bold">{movie.title}</h1>
        <p className="text-neutral-400 mt-2 text-center">{movie.tagline}</p>
        <p className="text-md text-neutral-200 mt-2">Release Date: {movie.release_date}</p>
        <p className="text-md text-neutral-200 mt-2">Runtime: {movie.runtime} minutes</p>
        <p className="text-md text-neutral-200 mt-2">Genres: {movie.genres?.map(genre => genre.name).join(', ')}</p>
        <div className="mt-4 text-lg">
          <p>Description: {movie.overview || "No additional description provided."}</p>
        </div>
      </div>
    </div>
  );
}
