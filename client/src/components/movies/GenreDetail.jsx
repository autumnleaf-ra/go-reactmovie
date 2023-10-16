import React from "react";
import { Link } from "react-router-dom";

const GenreDetail = ({ movies }) => {
  return (
    <ul>
      {Array.isArray(movies) ? (
        movies.map((movie) => (
          <li key={movie.id}>
            <Link to={`/movies/${movie.id}`}>{movie.title}</Link>
          </li>
        ))
      ) : (
        <p>Oops.. There's no movie data.</p>
      )}
    </ul>
  );
};

export default React.memo(GenreDetail);
