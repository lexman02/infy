import React, { useState } from 'react';

export default function OtherVideos({ otherVideos }) {
    const [expanded, setExpanded] = useState(false);
    const displayVideos = expanded ? otherVideos : otherVideos.slice(0, 5);


    function toggleExpand() {
        setExpanded(!expanded);
    }

    return (
        <div className="">
            <h2 className="text-2xl font-bold mb-4">More Videos</h2>
            <div className="flex space-x-4 overflow-x-auto scrollbar-hidden">
                {displayVideos.map(video => (
                    video.site === "YouTube" && (
                        <iframe
                            key={video.id}
                            src={`https://www.youtube.com/embed/${video.key}`}
                            title="YouTube video player"
                            frameBorder="0"
                            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                            allowFullScreen
                            className="w-full h-[250px] md:h-[450px] rounded-lg"
                        ></iframe>
                    )
                ))}
            </div>
            {otherVideos.length > 5 && (
                <button className="mt-4 w-full bg-indigo-800/70 text-white py-2 rounded" onClick={toggleExpand}>
                    {expanded ? "Show Less" : "Show More"}
                </button>
            )}
        </div>
    )
}