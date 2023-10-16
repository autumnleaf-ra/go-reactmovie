import axios from "axios";
import React, { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import GenreDetail from "../../components/movies/GenreDetail";

function ShowMoviesGenre() {
  let { id } = useParams();
  const [movies, setMovie] = useState([]);
  /* Set Loading */
  const [loaded, setLoaded] = useState(false);
  /* Error  */
  const [errorMessage, seterrorMessage] = useState(null);

  useEffect(() => {
    const fetchMovies = async () => {
      try {
        const result = await axios(`http://localhost:4000/genres/${id}/movies`);
        /* result , data , json movies */
        await setMovie(result.data.movies);
        setLoaded(true);
      } catch (err) {
        seterrorMessage(err.response.data);
      }
    };
    fetchMovies();
  }, [id]);

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
        <GenreDetail movies={movies} />
      )}
    </>
  );
}

export default ShowMoviesGenre;
