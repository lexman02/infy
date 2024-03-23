import React, { useState } from "react";

export default function Favorites() {
    const [movies, setMovies] = useState([]);
    const [inputValue, setInputValue] = useState([]);

    const addMovie = () => {
        if (inputValue) {
            setMovies([...movies, inputValue]);
            setInputValue('');
        }
    };

    const removeMovie = (index) => {
        const newMovies = movies.filter((_, i) => i !== index);
        setMovies(newMovies);
    };


    return (
        <div className="searchBar">
            <input
                type="text"
                value={inputValue}
                onChange={(e) => setInputValue(e.target.value)}
                placeholder="Search a movie!"
            />
            <button onClick={addMovie}>Add</button>
            <ul>
                {movies.map((movie, index) => (
                    <li key={index}>
                        {movie}
                        <button onClick={() => removeMovie(index)}>Remove</button>
                    </li>
                ))}
            </ul>
        </div>
    );
}