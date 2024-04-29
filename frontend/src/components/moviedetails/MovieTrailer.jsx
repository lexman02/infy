import React, { useState, useEffect } from "react"
import axios from "axios"

export default function MovieTrailer({ movieID, setOtherVideos, setErrors }) {
    const [mainTrailer, setMainTrailer] = useState(null);

    useEffect(() => {
        axios.get(`http://localhost:8000/movies/${movieID}/trailers`)
            .then(response => {
                const trailers = response.data.results;
                const main = trailers.find(trailer => trailer.type === "Trailer");
                setMainTrailer(main);

                setOtherVideos(trailers.filter(trailer => trailer.type !== "Trailer"));
            })
            .catch(error => {
                console.error(error);
                if (error.response && error.response.data && error.response.data.message) {
                    setErrors(errors => [...errors, error.response.data.message]);
                } else {
                    setErrors(errors => [...errors, "An error occurred while fetching trailers."]);
                }
            });
    }, [movieID, setOtherVideos, setErrors]);

    return (
        mainTrailer && (
            <div className="">
                <h2 className="text-2xl font-bold mb-4">Trailer</h2>
                {mainTrailer.site === "YouTube" && (
                    <iframe
                        src={`https://www.youtube.com/embed/${mainTrailer.key}`}
                        title="YouTube video player"
                        frameBorder="0"
                        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                        allowFullScreen
                        className="w-full h-[250px] md:h-[450px] rounded-lg"
                    ></iframe>
                )}
            </div>
        )
    )
}