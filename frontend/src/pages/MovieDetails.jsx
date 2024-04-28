import React, { useEffect, useState, useContext } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import { UserContext } from '../contexts/UserProvider';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function MovieDetails() {
  const { movieID } = useParams();
  const [movie, setMovie] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const { userData } = useContext(UserContext); 
  const [errorMessage, setErrorMessage] = useState('');

  const handleCloseError = () => {
        setErrorMessage('');
  };

  // Fetch movie details
  useEffect(() => {
    if (movieID) {
      setIsLoading(true);
      axios.get(`http://localhost:8000/movies/${movieID}`)
        .then(response => {
          setMovie(response.data);
          setIsLoading(false);
        })
        .catch(error => {
          console.error(error);
            setIsLoading(false);
            if (error.response && error.response.data && error.response.data.message) {
                // If the error contains a specific message, set that as the errorMessage
                setErrorMessage(error.response.data.message);
            } else {
                // If no specific message is available, set a generic error message
                setErrorMessage('An error occurred while fetching details, please wait then try again.');
            }
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

          <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
              <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                  {errorMessage}
              </Alert>
          </Snackbar>

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
