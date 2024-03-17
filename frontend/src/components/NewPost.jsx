import React, { useState } from 'react';
import MovieSearchPost from './MovieSearchPost';
import { PaperAirplaneIcon } from '@heroicons/react/20/solid';
import axios from 'axios';

export default function NewPost({ onNewPost }) {
    const [selectedMovie, setSelectedMovie] = useState(null);

    const handleSelectResult = (movie) => {
        setSelectedMovie(movie);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = Object.fromEntries(new FormData(e.target));
        data.movie_id = selectedMovie.id.toString();
        await axios.post('http://localhost:8000/posts/', data, {withCredentials: true})
            .then(() => {
                setSelectedMovie(null);
                onNewPost();
            })
            .catch(error => {
                console.error(error);
            });
    }

    return (
        <div>
            {!selectedMovie && <MovieSearchPost onSelectResult={handleSelectResult} />}
            {selectedMovie && (
                <div>
                    <div className="transition-all duration-500 cursor-pointer" onClick={() => setSelectedMovie(null)}>
                        <div className="flex hover:bg-black/30 hover:rounded-lg">
                            <img src={`https://image.tmdb.org/t/p/original/${selectedMovie.poster_path}`} alt={selectedMovie.title} className="w-20 h-32 object-cover rounded-lg" />
                            <div className="px-4 w-full">
                                <h2 className="text-xl font-bold text-neutral-200 mb-2">{selectedMovie.title}</h2>
                                <p className="text-neutral-400 line-clamp-2">{selectedMovie.overview}</p>
                            </div>
                        </div>
                    </div>

                    <form onSubmit={handleSubmit}>
                        <div className="mt-6 relative">
                            <label htmlFor="content" className="block mb-2 text-md font-medium"></label>
                            <textarea id="content" name="content" placeholder="Anything to say?" className="border rounded-lg border-indigo-500/30 bg-indigo-950/70 text-neutral-200 shadow-inner block w-full p-2.5" data-gramm="false" autoComplete="off"/>
                            {/* disable auto complete and grammarly */}
                            <button className="absolute bottom-2 right-2 bg-indigo-900 hover:bg-indigo-800 text-neutral-50 rounded-full p-2">
                                <PaperAirplaneIcon className="h-4 w-4" />
                            </button>
                        </div>
                    </form>
                </div>
            )}
        </div>
    );
}