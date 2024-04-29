import Navbar from './components/Navbar.jsx';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Watchlist from './pages/Watchlist.jsx';
import Profile from './pages/Profile.jsx';
import Search from './pages/Search.jsx'
import Home from './pages/Home.jsx';
import Login from "./pages/Login.jsx";
import Signup from "./pages/Signup.jsx";
import DetailedPost from "./pages/DetailedPost.jsx";
import logo from './img/logo.png';
import MovieDetails from './pages/MovieDetails.jsx';
import ActorDetails from './pages/ActorDetails.jsx';
import axios from 'axios';
import { useContext, useState } from 'react';
import { UserContext } from './contexts/UserProvider';
import Admin from './pages/Admin.jsx';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';
import NotFound from './pages/NotFound.jsx';

export default function App() {
  const { userData, setUserData } = useContext(UserContext);
  const [errorMessage, setErrorMessage] = useState('');

  const handleCloseError = () => {
    setErrorMessage('');
  };

  function logout() {
    axios.post('http://localhost:8000/auth/logout', {}, { withCredentials: true })
      .then(() => {
        setUserData(null);
        window.location.href = '/login';
      })
      .catch(error => {
        console.error(error);
        if (error.response && error.response.data && error.response.data.message) {
          // If the error contains a specific message, set that as the errorMessage
          setErrorMessage(error.response.data.message);
        } else {
          // If no specific message is available, set a generic error message
          setErrorMessage('An error occured while liking.');
        }
      });
  }

  function handleSignup() {
    window.location.href = '/signup';
  }

  return (
    <div className="bg-space-infy flex flex-col bg-black/10 min-h-screen text-neutral-100 font-sans">
      <div className="flex justify-end border-b border-neutral-700">
        {userData === null ? <button onClick={handleSignup} className="flex justify-center w-20 bg-violet-900/70 text-neutral-50 rounded-lg p-1 m-2 hover:bg-violet-950">Signup</button> : null}
        {userData !== null ? <button onClick={logout} className="flex justify-center w-20 bg-violet-900/70 text-neutral-50 rounded-lg p-1 m-2 hover:bg-violet-950">Logout</button> : null}
      </div>
      <div className="flex justify-center">
        <a href="/">
          <img src={logo} alt="Infy Logo" className="w-48 h-48" />
        </a>
      </div>
      <Router>
        <Routes>
          <Route exact path="/login" element={<Login />} />
          <Route exact path="/signup" element={<Signup />} />
          <Route exact path="/" element={<Home />} />
          <Route path="/post/:postID" element={<DetailedPost />} />
          <Route path="/profile/:username" element={<Profile />} />
          <Route exact path="/watchlist" element={<Watchlist />} />
          <Route exact path="/search" element={<Search />} />
          <Route exact path="/movie/:movieID" element={<MovieDetails />} />
          <Route exact path="/admin" element={<Admin />} />
          <Route exact path="/actor/:actorID" element={<ActorDetails />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
        <Navbar />
      </Router>

      <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
        <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
          {errorMessage}
        </Alert>
      </Snackbar>

    </div>
  )
}
