import { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

export default function SimilarMovies({ movieID, setErrors }) {
    const [similarMovies, setSimilarMovies] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        axios.get(`http://localhost:8000/movies/${movieID}/similar`)
            .then((response) => {
                if (response.data && response.data.results && Array.isArray(response.data.results)) {
                    const filteredMovies = response.data.results.filter((movie) => movie.poster_path !== null && movie.poster_path !== "");
                    setSimilarMovies(filteredMovies);
                } else {
                    setErrors((errors) => [...errors, "Failed to fetch similar movies."]);
                }
            })
            .catch(error => {
                if (error.response && error.response.data && error.response.data.message) {
                    setErrors((errors) => [...errors, error.response.data.message]);
                } else {
                    setErrors((errors) => [...errors, "An error occurred while fetching similar movies."]);
                }
            });
    }, [movieID, setErrors]);

    return (
        <div className="">
            <h2 className="text-2xl font-bold mb-4">Similar Movies</h2>
            <div className="flex overflow-x-auto scrollbar-hidden space-x-4">
                <div className="flex gap-4">
                    {similarMovies.map((simMovie, index) => (
                        <div key={index} className="w-[180px] md:w-[200px] lg:w-[220px] cursor-pointer" onClick={() => navigate(`/movie/${simMovie.id}`)}>
                            <img
                                className="w-full h-[270px] md:h-[300px] lg:h-[330px] object-cover rounded-lg"
                                src={`https://image.tmdb.org/t/p/original/${simMovie.poster_path}`}
                                alt={simMovie.title}
                                title={`Go to ${simMovie.title}'s details`}
                            />
                            <h3 className="text-center text-neutral-200 text-sm font-semibold mt-2 line-clamp-1">{simMovie.title}</h3>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    )
}