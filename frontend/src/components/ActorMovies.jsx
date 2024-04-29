import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

export default function ActorMovies({ actorID, setErrors }) {
    const [actorMovies, setActorMovies] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        axios.get(`http://localhost:8000/movies/actor/${actorID}/movies`)
            .then(response => {
                // Filter to include only movies with a valid poster path
                const filteredMovies = response.data.cast.filter(movie => movie.poster_path !== null && movie.poster_path !== '');
                setActorMovies(filteredMovies);
            })
            .catch(error => {
                if (error.response && error.response.data && error.response.data.message) {
                    setErrors(errors => [...errors, error.response.data.message]);
                } else {
                    setErrors(errors => [...errors, "Failed to fetch actor's movies"]);
                }
            });
    }, [actorID, setErrors]);

    function handleMovieClick(movieID) {
        // Redirect to the movie details page
        navigate(`/movie/${movieID}`);
    }

    return (
        <div className="w-full space-y-4">
            <h2 className="text-2xl font-bold text-center md:text-start">Actor's Movies</h2>
            <div className="flex flex-wrap justify-center md:justify-start gap-4">
                {actorMovies.map((movie, index) => (
                    <div key={index} className="shrink-0 w-[180px] md:w-[200px] lg:w-[220px] cursor-pointer" onClick={() => handleMovieClick(movie.id)}>
                        <img className="w-full h-[270px] md:h-[300px] lg:h-[330px] object-cover rounded-lg"
                            src={`https://image.tmdb.org/t/p/w500${movie.poster_path}`}
                            alt={movie.title}
                            title={`Go to ${movie.title}'s details`}
                        />
                        <h3 className="text-center text-neutral-200 text-sm font-semibold mt-2 line-clamp-1">{movie.title}</h3>
                    </div>
                ))}
            </div>
        </div>
    )
}