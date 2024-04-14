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
        await axios.post('http://localhost:8000/posts/', data, { withCredentials: true })
            .then(() => {
                setSelectedMovie(null);

                if (data.watched) {
                    axios.post(`http://localhost:8000/profile/movies/add/watched`, { 'movieId': data.movie_id }, { withCredentials: true })
                        .catch(error => {
                            console.error(error);
                        });
                }

                onNewPost();
            })
            .catch(error => {
                console.error(error);
            });


        onNewPost();
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
                            {/* disable auto complete and grammarly */}
                            <textarea id="content" name="content" placeholder="Anything to say?" className="border rounded-t-lg border-indigo-900 bg-indigo-950/70 text-neutral-200 shadow-inner block w-full p-2.5" data-gramm="false" autoComplete="off" />
                            <div className="border border-t-0 border-indigo-900 rounded-b-lg flex items-center justify-between p-2">
                                <label className="inline-flex items-center cursor-pointer">
                                    <input type="checkbox" name="watched" value="true" className="sr-only peer" defaultChecked />
                                    <div className="relative w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4 peer-focus:ring-indigo-300 dark:peer-focus:ring-indigo-800 dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-indigo-600"></div>
                                    <span className="ms-3 text-sm font-medium text-gray-900 dark:text-gray-300">Add movie to watched?</span>
                                </label>
                                <button className="bg-indigo-600 hover:bg-indigo-800 text-neutral-50 rounded-full p-2">
                                    <PaperAirplaneIcon className="h-4 w-4" />
                                </button>
                            </div>
                        </div>
                    </form>
                </div>
            )}
        </div>
    );
}