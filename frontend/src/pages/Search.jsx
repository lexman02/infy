import { useState, useEffect } from "react";
import axios from "axios";
import { Link } from "react-router-dom";

export default function Search() {
    const [inputValue, setInputValue] = useState('');
    const [searchResults, setSearchResults] = useState([]);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        if (inputValue !== '') {
            setIsLoading(true);
            setSearchResults([]);
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
            setIsLoading(true);
            setSearchResults([]);
            axios.get(`http://localhost:8000/movies/trending/day`)
                .then(response => {
                    setSearchResults(response.data.results);
                    setIsLoading(false);
                })
                .catch(error => {
                    console.error(error);
                    setIsLoading(false);
                });
        }
    }, [inputValue]);

    return (
        <div className="md:my-6 md:mx-60 flex-grow bg-indigo-950/70 rounded-lg p-4 space-y-4">
            <div>
                <label htmlFor="movie-search" className="block mb-2 text-md font-medium text-neutral-200">Discover the InfyVerse</label>
                <input type="text" id="movie-search" name="movie-search" placeholder="Search for a movie..." onChange={e => setInputValue(e.target.value)} className="border rounded-lg border-indigo-500/30 bg-indigo-950/70 text-neutral-200 sm:text-sm block w-full p-2.5" />
            </div>

            {isLoading ? (
                <div className="flex justify-center items-center h-2/3 p-6" >
                    {/* <div className="animate-spin rounded-full h-16 w-16 border-t-2 border-b-2 border-indigo-500"></div> */}
                    < svg className="animate-spin -ml-1 mr-3 h-8 w-8 text-indigo-500" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" >
                        <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                        <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                </div >
            ) : (
                <div class="grid grid-flow-row grid-cols-3 gap-x-2">
                    {searchResults.map((result, index) => (
                        <Link to={"/movie/" + result.id} key={result.id} className="cursor-pointer">
                            <div className="flex hover:bg-black/30 hover:rounded-lg">
                                <img src={`https://image.tmdb.org/t/p/original/${result.poster_path}`} alt={result.title} className="w-20 h-32 object-cover rounded-lg" />
                                <div className="px-4 w-full">
                                    <h2 className="text-xl font-bold text-neutral-200 mb-2">{result.title}</h2>
                                    <p className="text-neutral-400 line-clamp-2">{result.overview}</p>
                                </div>
                            </div>
                            {index !== searchResults.length - 1 && <hr className="border-indigo-500/30 my-2" />}
                        </Link>
                    ))}
                </div>
            )}
        </div >
    );
}