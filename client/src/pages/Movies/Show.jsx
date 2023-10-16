import axios from "axios";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import MovieDetail from "../../components/movies/MovieDetail";

function ShowMovie() {
  let { id } = useParams();
  const [movie, setMovie] = useState([]);
  /* Set Loading */
  const [loaded, setLoaded] = useState(false);
  /* Error  */
  const [errorMessage, seterrorMessage] = useState(null);

  useEffect(() => {
    const fetchMovies = async () => {
      try {
        const result = await axios(`http://localhost:4000/movies/${id}`);
        /* result , data , json movies */
        await setMovie(result.data.movie);
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
        <MovieDetail movie={movie} />
      )}
    </>
  );
}

export default ShowMovie;
