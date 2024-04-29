import { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

export default function MovieCast({ movieID, setErrors }) {
    const [cast, setCast] = useState([]);
    const [expanded, setExpanded] = useState(false);
    const visibleCast = expanded ? cast : cast.slice(0, 10);
    const navigate = useNavigate();


    useEffect(() => {
        axios.get(`http://localhost:8000/movies/${movieID}/cast`)
            .then((response) => {
                if (response.data && response.data.cast && Array.isArray(response.data.cast)) {
                    const filteredCast = response.data.cast.filter((actor) => actor.profile_path !== null && actor.profile_path !== "");
                    setCast(filteredCast);
                } else {
                    setErrors((errors) => [...errors, "Failed to fetch cast."]);
                }
            })
            .catch(error => {
                if (error.response && error.response.data && error.response.data.message) {
                    setErrors((errors) => [...errors, error.response.data.message]);
                } else {
                    setErrors((errors) => [...errors, "An error occurred while fetching cast."]);
                }
            });
    }, [movieID, setErrors]);

    function handleActorClick(actorName, actorProfilePath) {
        const params = `name=${encodeURIComponent(actorName).replace(/%20/g, '%2B')}`;
        axios.get(`http://localhost:8000/people/search?${params}`)
            .then(response => {
                const results = response.data.results;
                const actor = results.find(a => a.profile_path === actorProfilePath);

                if (response.data && actor) {
                    navigate(`/actor/${actor.id}`);
                } else {
                    setErrors((errors) => [...errors, "No matching actor found with the given image"]);
                }
            })
            .catch(error => {
                if (error.response && error.response.data && error.response.data.message) {
                    setErrors((errors) => [...errors, error.response.data.message]);
                } else {
                    setErrors((errors) => [...errors, "An error occurred while fetching actor details."]);
                }
            });
    }

    function toggleExpand() {
        setExpanded(!expanded);
    }

    return (
        <div>
            <h2 className="text-2xl font-bold mb-4">Cast</h2>
            <div className="grid grid-cols-2 md:grid-cols-5 gap-4 overflow-x-auto scrollbar-hidden">
                {visibleCast.map((actor, index) => (
                    <div key={index} className="bg-indigo-900/40 p-2 rounded-xl shadow hover:shadow-lg transition-shadow duration-300 cursor-pointer flex flex-col items-center justify-start" onClick={() => handleActorClick(actor.name, actor.profile_path)}>
                        <img src={`https://image.tmdb.org/t/p/w500${actor.profile_path}`} alt={actor.name} className="rounded-lg mb-2 object-cover h-full w-auto" />
                        <p className="font-semibold text-sm text-center">{actor.name}</p>
                        <p className="text-neutral-400 text-xs text-center">{`as ${actor.character}`}</p>
                    </div>
                ))}
            </div>
            {cast.length > 10 && (
                <button className="mt-4 w-full bg-indigo-800/70 text-white py-2 rounded" onClick={toggleExpand}>
                    {expanded ? "Show Less" : "Show More"}
                </button>
            )}
        </div>
    )
}