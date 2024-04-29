import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';
import ActorMovies from '../components/ActorMovies';

export default function ActorDetails() {
    const { actorID } = useParams();
    const [actor, setActor] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const [errors, setErrors] = useState([]);
    const navigate = useNavigate();

    const handleCloseError = () => {
        setErrors([]);
    };

    useEffect(() => {
        if (actorID) {
            setIsLoading(true);
            axios.get(`http://localhost:8000/movies/actor/${actorID}`)
                .then(response => {
                    setActor(response.data);
                    setIsLoading(false);
                })
                .catch(error => {
                    if (error.response && error.response.data && error.response.data.message) {
                        setErrors(errors => [...errors, error.response.data.message]);
                    } else {
                        setErrors(errors => [...errors, "Failed to fetch actor details."]);
                    }
                    setIsLoading(false);
                });
        }
    }, [actorID]);

    function showError() {
        return errors.length > 0;
    }

    if (isLoading) {
        return <div>Loading...</div>;
    }

    if (!actor) {
        navigate('/404');
    }

    return (
        <div className="md:my-6 md:mx-60 flex-grow p-4 md:p-0 space-y-8">
            <div className="flex flex-col md:flex-row space-y-4 md:space-x-8 lg:items-start items-center">
                {/* Actor Image Section */}
                <div className="flex-shrink-0">
                    <img
                        src={`https://image.tmdb.org/t/p/original/${actor.profile_path || "default_placeholder.jpg"}`}
                        alt={actor.name || "Unknown"}
                        className="rounded-lg mb-4 w-full lg:max-w-xs object-cover"
                    />
                </div>

                {/* Actor Details Section */}
                <div className="flex flex-col text-center md:text-start items-center lg:items-start space-y-4 md:space-y-2">
                    <h1 className="text-4xl font-extrabold">
                        {actor.name || "Name not available"}
                    </h1>
                    <div className="flex bg-indigo-900/70 rounded-lg py-1 px-3 text-neutral-300 font-semibold">
                        <p className="text-xl">
                            {new Date(actor.birthday).toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' })}
                        </p>
                        <span className="text-xl mx-2">•</span>
                        <p className="text-xl">
                            {actor.place_of_birth || "Place of birth not available"}
                        </p>
                    </div>
                    <div className="text-lg font-medium text-neutral-50">
                        <p>{actor.biography || "No biography available."}</p>
                    </div>
                </div>
            </div>

            {/* Actor's Movies Section */}
            <ActorMovies actorID={actorID} setErrors={setErrors} />

            {
                errors.map((error, index) => (
                    <Snackbar key={index} open={showError()} autoHideDuration={6000} onClose={handleCloseError}>
                        <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                            {error}
                        </Alert>
                    </Snackbar>
                ))
            }
        </div>
    );
}