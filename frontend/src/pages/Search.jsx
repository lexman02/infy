import { useState, useEffect } from "react";
import axios from "axios";
import MovieSearch from "../components/MovieSearch";


export default function Search(){
    const [selectedMovie, setSelectedMovie] = useState(null);

    const handleSelectResult = (movie) => {
        setSelectedMovie(movie);
    };

    return (
        <div className="md:my-6 md:mx-60 p-3 flex-grow bg-black/40">
                <div>
                    {!selectedMovie && <MovieSearch onSelectResult={handleSelectResult} />}
                </div>
        </div>
    );
} 