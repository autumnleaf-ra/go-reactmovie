/* third party */
import { Link } from "react-router-dom";
import React, { useEffect, useState } from "react";
import axios from "axios";

const MovieList = () => {
  const [movies, setMovies] = useState([]);
  /* Set Loading */
  const [loaded, setLoaded] = useState(false);
  /* Error  */
  const [errorMessage, seterrorMessage] = useState(null);

  useEffect(() => {
    const fetchMovies = async () => {
      try {
        const result = await axios(`http://localhost:4000/movies`);
        /* result , data , json movies */
        await setMovies(result.data.movies);
        setLoaded(true);
      } catch (err) {
        seterrorMessage(err.response.data);
      }
    };
    fetchMovies();
  }, []);

  return (
    <>
      {!loaded ? (
        (() => {
          if (errorMessage) {
            return (
              <div className="row">
                <p>Oopss ... {errorMessage}</p>
              </div>
            );
          } else {
            <div className="row">
              <p>Loading ...</p>
            </div>;
          }
        })()
      ) : (
        <div className="row">
          {movies.map((movie, index) => (
            <div className="col-sm-4 mb-2" key={index}>
              <div className="card">
                <div className="card-body">
                  <h5 className="card-title">{movie.title}</h5>
                  <p className="card-text">{movie.description}</p>
                  <Link to={`/movies/${movie.id}`} className="btn btn-primary">
                    Go somewhere
                  </Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </>
  );
};

export default MovieList;
