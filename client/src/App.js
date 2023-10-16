/* third party */
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";

/*  Internal Source */
import Menu from "./components/Menu";
import Home from "./pages/Home";
import Movies from "./pages/Movies";
import ShowMovie from "./pages/Movies/Show";
import Genres from "./pages/Genres";
import Admin from "./pages/Admin";

/* style */
import "./App.css";
import ShowMoviesGenre from "./pages/Genres/Show";
import MovieForm from "./components/movies/MovieForm";

function App() {
  return (
    <Router>
      <div className="container">
        <div className="row">
          <h1 className="mt-3 text-center">Go React Movie Project</h1>
          <hr className="mb-3" />
        </div>
        <div className="row">
          <div className="col-sm-2 mb-3">
            <Menu />
          </div>
          <div className="col-10">
            <Routes>
              <Route path="/" element={<Home />} />
              <Route exact path="/movies/:id" element={<ShowMovie />} />
              <Route path="/movies" element={<Movies />} />
              <Route
                exact
                path="/genres/:id/movies"
                element={<ShowMoviesGenre />}
              />
              <Route path="/genres" element={<Genres />} />
              <Route path="/admin" element={<Admin />} />
              <Route
                exact
                path="/admin/movies/create"
                element={<MovieForm />}
              />
              <Route
                exact
                path="/admin/movies/:id/edit"
                element={<MovieForm />}
              />
            </Routes>
          </div>
        </div>
      </div>
    </Router>
  );
}

export default App;
