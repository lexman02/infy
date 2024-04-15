import React, { useState, useEffect } from "react";
import axios from "axios";
import { useParams } from 'react-router-dom';
import AddToWatchlist from "../components/AddToWatchlist";

export default function Favorites() {
    return (
        <div>
            <AddToWatchlist movieID={697} />
        </div>
    );
}