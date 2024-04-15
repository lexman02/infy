import React, { useState } from 'react';
import axios from 'axios';

export default function AddToWatchlist({ isAdded = false, movieID }) {
    const [addded, setAdded] = useState(isAdded);

    // Add to list
    async function addMovie() {
        await axios.post(`http://localhost:8000/profile/movies/add/watchlist`, { 'movieId': movieID.toString() }, { withCredentials: true })
            .then(() => {
                setAdded(true);
            })
            .catch((error) => {
                console.error('ERROR adding to watchlist: ', error);
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
                console.error('ERROR removing from watchlist: ', error);
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
        </div>
    )
}