import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export default function NotFound() {
    const [count, setCount] = useState(5);
    const [randomGif] = useState(() => {
        const gifs = [
            "https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExOGVmanlia3poYnFnZHdhanA2d21xbWZwcWFmbjNnZ2pxamliYmYyeiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/9JgjpeyN48rdtzIvOZ/giphy.gif",
            "https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExd3Q5ZnowanNkZTkyY291YjVjd2Q4Z3p5eWkyNGEydzlja2hoNnU5NCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/1EmBoG0IL50VIJLWTs/giphy.gif",
            "https://media.giphy.com/media/3o7aTskHEUdgCQAXde/giphy.gif",
        ]
        return gifs[Math.floor(Math.random() * gifs.length)];
    });
    const navigate = useNavigate();

    useEffect(() => {
        const interval = setInterval(() => {
            if (count > 0) {
                setCount(count - 1);
            } else {
                clearInterval(interval);
                navigate('/');
            }
        }, 1000);

        return () => clearInterval(interval);
    });

    return (
        <div className="flex-grow flex flex-col justify-evenly items-center">
            <div className="flex flex-col justify-center items-center">
                <h1 className="text-6xl font-extrabold text-indigo-500/60">404</h1>
                <p className="text-xl font-medium text-indigo-200/80">Looks like you got lost in the infy-verse!</p>
            </div >
            <img src={randomGif} alt="404 GIF" className="max-w-1/4 max-h-1/4 md:rounded-lg" />
            <h2 className="text-2xl font-medium text-indigo-300/50">Redirecting you in {count}...</h2>
        </div >
    );
}