/* third party */
import { Link } from "react-router-dom";
import React, { useEffect, useState } from "react";
import axios from "axios";

const MovieTable = () => {
  const [movies, setMovies] = useState([]);
  /* Set Loading */
  const [loaded, setLoaded] = useState(false);
  /* Error  */
  const [errorMessage, seterrorMessage] = useState(null);

  const confirmDelete = async (id) => {
    const payload = {
      id: id.toString(),
    };

    await axios.post(
      "http://localhost:4000/admin/movies/delete",
      JSON.stringify(payload)
    );
    setMovies([]);
    fetchMovies();
  };

  const fetchMovies = async () => {
    try {
      const result = await axios(`http://localhost:4000/movies`);
      if (result.data.movies !== null) {
        /* result , data , json movies */
        await setMovies(result.data.movies);
        setLoaded(true);
      } else {
        seterrorMessage("data not found!");
      }
    } catch (err) {
      seterrorMessage(err.response.data);
    }
  };

  useEffect(() => {
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
        <>
          <div className="row">
            <div className="col-12">
              <Link to={"/admin/movies/create"} className="btn btn-primary">
                Add
              </Link>
            </div>
          </div>
          <div className="row">
            <div className="col-12">
              <table className="table">
                <thead>
                  <tr>
                    <td>No</td>
                    <td>Name</td>
                    <td></td>
                  </tr>
                </thead>
                <tbody>
                  {movies.map((movie, index) => (
                    <tr key={index}>
                      <td>{index + 1}</td>
                      <td>
                        <Link to={`/movies/${movie.id}`}>{movie.title}</Link>
                      </td>
                      <td>
                        <div className="btn-group">
                          <button
                            className="btn btn-secondary btn-sm dropdown-toogle rounded"
                            type="button"
                            data-bs-toggle="dropdown"
                            aria-expanded="false"
                          >
                            Action
                          </button>
                          <ul className="dropdown-menu">
                            <li>
                              <span className="dropdown-item">
                                <Link to={`/admin/movies/${movie.id}/edit`}>
                                  Edit
                                </Link>
                              </span>
                            </li>
                            <li>
                              <span
                                className="dropdown-item"
                                style={{ cursor: "pointer" }}
                                onClick={() => {
                                  if (window.confirm("Are you sure ?")) {
                                    confirmDelete(movie.id);
                                  }
                                }}
                              >
                                Delete
                              </span>
                            </li>
                          </ul>
                        </div>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        </>
      )}
    </>
  );
};

export default MovieTable;
