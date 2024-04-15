import React, { useState, useEffect } from 'react';
import axios from 'axios';


export default function MovieSearch({ onSelectResult }){
    const [inputValue, setInputValue] = useState('');
    const [searchResults, setSearchResults] = useState([]);
    const [isLoading, setIsLoading] = useState(false);
    const inputClass = isLoading || searchResults.length > 0 ? 'border-b rounded-t-lg' : 'border rounded-lg';

    useEffect(() => {
        if (inputValue !== '') {
            setIsLoading(true);
            const fetchData = setTimeout(() => {
                const params = `title=${encodeURIComponent(inputValue).replace(/%20/g, '%2B')}`;
                axios.get(`http://localhost:8000/movies/search?${params}`)
                    .then(response => {
                        setSearchResults(response.data.results);
                        setIsLoading(false);
                    })
                    .catch(error => {
                        console.error(error);
                        setIsLoading(false);
                    });
            }, 1500);

            return () => clearTimeout(fetchData);
        } else {
            setSearchResults([]);
        }
    }, [inputValue]);

    return (
        <div>
            <div>
                <label htmlFor="movie-search" className="block mb-2 text-md font-medium text-neutral-200">Create a new post</label>
                <input type="text" id="movie-search" name="movie-search" placeholder="Search for a movie..." onChange={e => setInputValue(e.target.value)} className={`${inputClass} border-indigo-500/30 bg-indigo-950/70 text-neutral-200 sm:text-sm block w-full p-2.5`}/>
            </div>
            <div className="last:rounded-b-lg px-6 bg-indigo-950/70 shadow">
                {isLoading ? (
                    <div className="flex justify-center items-center h-2/3 p-6">
                        {/* <div className="animate-spin rounded-full h-16 w-16 border-t-2 border-b-2 border-indigo-500"></div> */}
                        <svg className="animate-spin -ml-1 mr-3 h-8 w-8 text-indigo-500" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                            <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                            <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                    </div>
                ) : searchResults.map((result, index) => (
                    <div key={result.id} className="cursor-pointer last:pb-6 first:pt-6" onClick={() => onSelectResult(result)}>
                        <div className="flex flex-col">
                            <div className="flex hover:bg-black/30 hover:rounded-lg">
                                <img src={`https://image.tmdb.org/t/p/original/${result.poster_path}`} alt={result.title} className="w-20 h-32 object-cover rounded-lg" />
                                <div className="px-4 w-full">
                                    <h2 className="text-xl font-bold text-neutral-200 mb-2">{result.title}</h2>
                                    <p className="text-neutral-400 line-clamp-2">{result.overview}</p>
                                </div>
                            </div>
                            {index !== searchResults.length - 1 && <hr className="border-indigo-500/30 my-6"/>}
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}