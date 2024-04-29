import React, { useState, useEffect } from 'react';
import axios from 'axios';

export default function MovieReviews({ movieID, setErrors }) {
    const [reviews, setReviews] = useState([]);
    const [expandedReviews, setExpandedReviews] = useState(false);
    const visibleReviews = expandedReviews ? reviews : reviews.slice(0, 5);

    useEffect(() => {
        axios.get(`http://localhost:8000/movies/${movieID}/reviews`)
            .then(response => {
                if (response.data && Array.isArray(response.data.results)) {
                    setReviews(response.data.results);
                } else {
                    setErrors(errors => [...errors, 'Failed to fetch reviews.']);
                }
            })
            .catch(error => {
                if (error.response && error.response.data && error.response.data.message) {
                    setErrors(errors => [...errors, error.response.data.message]);
                } else {
                    setErrors(errors => [...errors, 'An error occurred while fetching reviews.']);
                }
            });
    }, [movieID, setErrors]);

    function toggleExpandReviews() {
        setExpandedReviews(!expandedReviews);
    }

    return (
        reviews.length > 0 && (
            <div className="">
                <h2 className="text-2xl font-bold mb-4">Reviews</h2>
                <div className="space-y-2">
                    {visibleReviews.map((review, index) => (
                        <div key={index} className="review-card bg-neutral-950/50 p-6 rounded-lg shadow-lg">
                            <p className="text-lg text-neutral-100 leading-relaxed">{review.content}</p>
                            <div className="mt-4 border-t border-neutral-700 pt-4">
                                <p className="text-neutral-400 italic">- {review.author}</p>
                            </div>
                        </div>
                    ))}
                </div>
                {reviews.length > 5 && (
                    <button className="mt-4 w-full bg-indigo-800/70 text-white py-2 rounded" onClick={toggleExpandReviews}>
                        {expandedReviews ? "Show Less" : "Show More"}
                    </button>
                )}
            </div>
        )
    );
}