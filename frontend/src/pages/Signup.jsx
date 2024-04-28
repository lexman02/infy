import React from 'react';
import axios from 'axios';
import { useState } from 'react';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function Signup() {
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = Object.fromEntries(new FormData(e.target));
        try {
            await axios.post('http://localhost:8000/auth/signup', data, {withCredentials: true});
            // Redirect the user to the profile page
            window.location.href = '/profile';
        } catch (error) {
            console.error(error);
            if (error.response && error.response.data && error.response.data.message) {
                // If the error contains a specific message, set that as the errorMessage
                setErrorMessage(error.response.data.message);
            } else {
                // If no specific message is available, set a generic error message
                setErrorMessage('An error occurred while signing up, please check required fields and try again.');
            }
          
        }
    }

    return (
        <div className="bg-neutral-900 rounded-lg w-80 h-160 mx-auto text-center">
            <div className="text-neutral-50 flex justify-center content-center align-middle">
                <h1>Join the InfyVerse!</h1>
            </div>
            <br/>
            <div className="flex justify-center align-middle">
                <form onSubmit={handleSubmit}>
                    {/*Username field */}
                    <label htmlFor="username" className="text-neutral-50">Username:</label><br/>
                    <input type="text" id="username" name="username" className="hover:bg-neutral-400"></input><br/><br/>

                    {/*Password field */}
                    <label htmlFor="password" className="text-neutral-50">Password:</label><br/>
                    <input type="password" id="password" name="password" className="hover:bg-neutral-400"></input><br/><br/>

                    {/*Confirm Password field */}
                    <label htmlFor="confirm_password" className="text-neutral-50">Confirm Password:</label><br/>
                    <input type="password" id="confirm_password" name="confirm_password" className="hover:bg-neutral-400"></input><br/><br/>

                    {/*Email field */}
                    <label htmlFor="email" className="text-neutral-50">Email:</label><br/>
                    <input type="text" id="email" name="email" className="hover:bg-neutral-400"></input><br/><br/>

                    {/*First Name field */}
                    <label htmlFor="first_name" className='text-neutral-50'>First Name:</label><br/>
                    <input type="text" id="first_name" name="first_name" className='hover:bg-neutral-400'></input><br/><br/>

                    {/*Last Name field */}
                    <label htmlFor="last_name" className='text-neutral-50'>Last Name:</label><br/>
                    <input type="text" id="last_name" name="last_name" className='hover:bg-neutral-400'></input><br/><br/>

                    {/*Date of Birth field*/}
                    <label htmlFor="date_of_birth" className='text-neutral-50'>Date of Birth:</label><br/>
                    <input type="date" id="date_of_birth" name="date_of_birth" className='hover:bg-neutral-400'></input><br/><br/>

                    {/*Submit button */}
                    <input type="submit" value="Submit" className="bg-violet-900 text-neutral-50 px-4 py-2 rounded-lg hover:bg-violet-950" />
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