import React, { useEffect, useState, useContext } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";
import { UserContext } from "../contexts/UserProvider";
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';
import MovieTrailer from "../components/moviedetails/MovieTrailer";
import MovieCast from "../components/moviedetails/MovieCast";
import SimilarMovies from "../components/moviedetails/SimilarMovies";
import MovieReviews from "../components/moviedetails/MovieReviews";
import OtherVideos from "../components/moviedetails/OtherVideos";
import AddToWatchlist from "../components/AddToWatchlist";

export default function MovieDetails() {
    const { movieID } = useParams();
    const [movie, setMovie] = useState(null);
    const [otherVideos, setOtherVideos] = useState([]);
    const [isLoading, setIsLoading] = useState(true);
    const [errors, setErrors] = useState([]);
    const { userData } = useContext(UserContext);
    const navigate = useNavigate();

    const handleCloseError = () => {
        setErrors([]);
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
                    if (error.response && error.response.data && error.response.data.message) {
                        setErrors(errors => [...errors, error.response.data.message]);
                    }
                    setIsLoading(false);
                });
        }
    }, [movieID]);

    function showError() {
        return errors.length > 0;
    }

    const isAdded = userData && userData.user.profile && userData.user.profile.preferences.watchlist.includes(movieID);
    console.log(isAdded,);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    if (!movie) {
        navigate('/404');
    }

    return (
        <div className="md:my-6 md:mx-60 flex-grow p-4 md:p-0 space-y-8">
            <div className="flex flex-col md:flex-row space-y-4 md:space-x-8 lg:items-start items-center">
                {/* Poster Section */}
                <div className="flex-shrink-0">
                    <img src={`https://image.tmdb.org/t/p/original/${movie.poster_path}`} alt={movie.title} className="rounded-lg mb-4 w-full lg:w-auto lg:max-w-xs object-cover" />
                </div>

                {/* Movie Details Section */}
                <div className="flex flex-col text-center md:text-start items-center lg:items-start space-y-4 md:space-y-2">
                    <div>
                        <h1 className="text-4xl font-extrabold">{movie.title}</h1>
                        <p className="text-neutral-400 text-2xl font-medium mt-2 md:m-0">{movie.tagline}</p>
                    </div>
                    <div className="flex bg-indigo-900/70 rounded-lg py-1 px-3 font-semibold">
                        <p className="text-xl text-neutral-200">
                            {new Date(movie.release_date).toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' })}
                        </p>
                        <span className="text-xl text-neutral-200 mx-2">•</span>
                        <p className="text-xl text-neutral-200">
                            {Math.floor(movie.runtime / 60)} hour(s) {movie.runtime % 60} minutes
                        </p>
                    </div>
                    <p className="text-lg text-neutral-50">
                        {movie.overview || "No additional description provided."}
                    </p>
                    <div className="flex space-x-4">
                        <AddToWatchlist isAdded={isAdded} movieID={movieID} />
                    </div>
                </div>
            </div>
            {/* Main Trailer Section */}
            <MovieTrailer movieID={movieID} setOtherVideos={setOtherVideos} setErrors={setErrors} />
            {/* Cast Section */}
            <MovieCast movieID={movieID} setErrors={setErrors} />
            {/* Similar Movies Section */}
            <SimilarMovies movieID={movieID} setErrors={setErrors} />
            {/* Reviews Section */}
            <MovieReviews movieID={movieID} setErrors={setErrors} />
            {/* Other Videos Section */}
            {otherVideos.length > 0 && (
                <OtherVideos otherVideos={otherVideos} />
            )}
            {
                errors.map((error, index) => (
                    <Snackbar key={index} open={showError()} autoHideDuration={6000} onClose={handleCloseError}>
                        <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                            {error}
                        </Alert>
                    </Snackbar>
                ))
            }
        </div >
    );
}
