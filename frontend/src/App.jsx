import Navbar from './components/Navbar.jsx';
import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Favorites from './pages/favorites.jsx';
import Profile from './pages/profile.jsx';
import Search from './pages/search.jsx'
import Home from './pages/Home.jsx';

export default function App() {
  return (
    <>
        <h1 className="text-3xl font-bold underline">
            Infy
        </h1>
        <Router>
            <Navbar />
            <Routes>
              <Route exact path="/" element={<Home/>}/>
              <Route exact path="/profile" element={<Profile/>}/>
              <Route exact path="/favorites" element={<Favorites/>}/>
              <Route exact path="/search" element={<Search/>}/>
            </Routes>
        </Router>
    </>
  )
}
