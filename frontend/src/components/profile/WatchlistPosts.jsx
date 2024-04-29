import React from "react"
import axios from "axios"
import { useState, useEffect } from "react"
import { Link } from "react-router-dom"
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function WatchlistPosts(watchlist) {
    const [movieData, setMovieData] = useState([])
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    useEffect(() => {
        if (!watchlist.watchlist) {
            return;
        }

        watchlist.watchlist.forEach(movieID => {
            axios.get(`http://localhost:8000/movies/${movieID}`)
                .then(response => {
                    setMovieData(prevData => {
                        const newData = [...prevData, response.data];
                        // Remove duplicates by filtering out movies with duplicate IDs
                        const uniqueData = newData.filter((movie, index, self) =>
                            index === self.findIndex(m => m.id === movie.id)
                        );
                        return uniqueData;
                    });
                })
                .catch(() => {
                    setErrorMessage('An error occurred while loading watchlist.');
                });
        });
    }, [watchlist]);

    return (
        <div className="bg-black/40 rounded-b-lg p-4">
            <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>

            {movieData.length > 0 ? (
                <div className="flex justify-start flex-wrap gap-4">
                    {(
                        movieData.map(movie => (
                            <Link to={`/movie/${movie.id}`} key={movie.id}>
                                <img src={`https://image.tmdb.org/t/p/original/${movie.poster_path}`} alt={movie.title} className=" w-24 h-36 object-cover rounded-lg" />
                            </Link>
                        ))
                    )}
                </div>
            ) : (
                <div className="p-4 text-neutral-300 font-medium text-lg rounded-b-lg text-center">
                    <h1 className="">No movies in the watchlist yet...</h1>
                </div>
            )}
        </div>
    )
}