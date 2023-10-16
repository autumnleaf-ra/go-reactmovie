/* third party */
import { Link } from "react-router-dom";
import axios from "axios";
import { useEffect, useState } from "react";

function GenreList() {
  const [genres, setGenres] = useState([]);
  /* Set Loading */
  const [loaded, setLoaded] = useState(false);
  /* Error  */
  const [errorMessage, seterrorMessage] = useState(null);

  const fetchGenres = async () => {
    try {
      const result = await axios(`http://localhost:4000/genres`);
      /* result , data , json movies */
      await setGenres(result.data.genres);
      setLoaded(true);
    } catch (err) {
      seterrorMessage(err.response.data);
    }
  };

  useEffect(() => {
    fetchGenres();
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
          {genres.map((genre, index) => (
            <div className="col-sm-2 mb-3" key={index}>
              <div className="card">
                <div className="card-body text-center">
                  <Link to={`/genres/${genre.id}/movies`}>
                    {genre.genre_name}
                  </Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </>
  );
}

export default GenreList;
