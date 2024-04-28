import React, { useState } from 'react';
import axios from 'axios';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function AddToWatchlist({ isAdded = false, movieID }) {
    const [addded, setAdded] = useState(isAdded);
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    // Add to list
    async function addMovie() {
        await axios.post(`http://localhost:8000/profile/movies/add/watchlist`, { 'movieId': movieID.toString() }, { withCredentials: true })
            .then(() => {
                setAdded(true);
            })
            .catch((error) => {
                console.error(error);
                if (error.response && error.response.data && error.response.data.message) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.message);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('ERROR occured adding to watchlist.');
                }
            });
    }

    // Remove from list
    async function removeMovie() {
        await axios.delete(`http://localhost:8000/profile/movies/watchlist/${movieID.toString}`, { withCredentials: true })
            .then((response) => {
                console.log(response.data.message);
                setAdded(false);
            })
            .catch((error) => {
                console.error(error);
                if (error.response && error.response.data && error.response.data.message) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.message);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('ERROR occured adding to watchlist.');
                }
            });
    }

    return (
        <div>
            {/* Add the movie to watchlist or remove if its already there */}
            {addded ? (
                <button onClick={removeMovie} className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">
                    Remove from Watchlist
                </button>
            ) : (
                <button onClick={addMovie} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                    Add to Watchlist
                </button>
            )}

            <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>

        </div>
    )
}