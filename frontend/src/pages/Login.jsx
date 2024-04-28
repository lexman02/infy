import React, {useContext, useState, useEffect} from 'react';
import axios from 'axios';
import { UserContext } from '../contexts/UserProvider';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function Login(){
    const {userData} = useContext(UserContext);
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    // Send user to profile once logged in
    useEffect(() => {
        if (userData !== null) {
            window.location.href = '/profile';
        }
    });

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = Object.fromEntries(new FormData(e.target));
        try {
            await axios.post('http://localhost:8000/auth/login', data, { withCredentials: true });
            // Redirect the user to the home page
            window.location.href = '/profile';
        } catch (error) {
            console.error(error);
            if (error.response && error.response.data && error.response.data.message) {
                // If the error contains a specific message, set that as the errorMessage
                setErrorMessage(error.response.data.message);
            } else {
                // If no specific message is available, set a generic error message
                setErrorMessage('An error occurred while logging in. Please check credentials and try again.');
            }
        }
    };

    return (
        <div className="bg-neutral-900 rounded-lg w-80 h-80 mx-auto text-center">
            <div className="text-neutral-50 flex justify-center content-center align-middle">
                <h1>Welcome back InfyNaut!</h1>
            </div>
            <br/>
            <div className="flex justify-center content-center align-middle">
                <form onSubmit={handleSubmit}>

                    {/*Email field */}
                    <label htmlFor="email" className="text-neutral-50">Email:</label><br/>
                    <input type="text" id="email" name="email" className="text-neutral-900"></input><br/><br/>

                    {/*Password field */}
                    <label htmlFor="password" className="text-neutral-50">Password:</label><br/>
                    <input type="password" id="password" name="password" className="text-neutral-900"></input><br/><br/>

                    {/*Submit field */}
                    <input type="submit" value="Log In" className="bg-violet-900 text-neutral-50 rounded-lg px-4 py-2 hover:bg-violet-950" />
                </form>
            </div>


            <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>

        </div>
    )
}