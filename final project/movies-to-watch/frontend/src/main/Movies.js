import { useState, useEffect } from "react";
import axios from 'axios';
import RightArrowIcon from '../assets/right-arrow-icon.png';
import EditIcon from '../assets/edit-icon.png';
import DeleteIcon from '../assets/delete-icon.png';

function Movies(props) {
    const [searchKeyword, setSearchKeyword] = useState("");
    const [movies, setMovies] = useState([]);
    const [filteredMovies, setFilteredMovies] = useState([]);

    const [showAddMovieModal, setShowAddMovieModal] = useState(false);
    const [title, setTitle] = useState("");
    const [isWatched, setIsWatched] = useState(null);

    const [isEditMovie, setIsEditMovie] = useState(false);
    const [selectedMovie, setSelectedMovie] = useState(null);

    useEffect(() => {
        axios.get(`http://localhost:8080/movie-list/${props.user.Id}`)
        .then(response => {
            // console.log(response.data);

            setMovies(response.data);
            setFilteredMovies(response.data);
        })
        .catch(error => {
            alert("Error: An unexpected error occurred. Please try again later.");
            console.log(error);
        });
    }, []);

    const filterMovies = () => {
        let filteredMovies = movies.filter(movie => movie.Title.toLowerCase().includes(searchKeyword.toLowerCase()));
        if (searchKeyword === "") {
            filteredMovies = [...movies];
        }
        setFilteredMovies(filteredMovies);
    }

    const addMovie = () => {
        if (title !== "" && isWatched !== null) {
            const movie = {
                Title: title,
                IsWatched: isWatched
            };

            axios.post(`http://localhost:8080/movie-list/${props.user.Id}`, movie)
            .then(response => {
                // console.log(response.data);

                setMovies(response.data);
                setFilteredMovies(response.data);
                setSearchKeyword("");

                setTimeout(function() {
                    alert("Success: Movie has been added in your list successfully.");
                }, 300);
                
                resetAddMovieForm();
            })
            .catch(error => {
                alert("Error: An unexpected error occurred. Please try again later.");
                console.log(error);
            });
        } else {
            alert("Error: Please fill out all fields.");
        }
    }

    const editMovie = (movie) => {
        if (!isEditMovie) {
            movie["IsWatched"] = !movie.IsWatched;
        } else {
            movie["Title"] = title
            movie["IsWatched"] = isWatched
        }

        axios.put(`http://localhost:8080/movie-list/edit/${movie.Id}`, movie)
        .then(response => {
            // console.log(response.data);

            setMovies(response.data);
            setFilteredMovies(response.data);
            setSearchKeyword("");

            if (isEditMovie) {
                setTimeout(function() {
                    alert("Success: Movie has been edited successfully.");
                }, 300);
            }
            
            resetAddMovieForm();
        })
        .catch(error => {
            alert("Error: An unexpected error occurred. Please try again later.");
            console.log(error);
        });
    }

    const deleteMovie = (movieId) => {
        axios.delete(`http://localhost:8080/movie-list/delete/${movieId}`)
        .then(response => {
            // console.log(response.data);

            setMovies(response.data);
            setFilteredMovies(response.data);
            setSearchKeyword("");

            setTimeout(function() {
                alert("Success: Movie has been deleted from your list successfully.");
            }, 300);
        })
        .catch(error => {
            alert("Error: An unexpected error occurred. Please try again later.");
            console.log(error);
        });
    }

    const resetAddMovieForm = () => {
        setShowAddMovieModal(false);
        setTitle("");
        setIsWatched(null);
        setIsEditMovie(false);
        setSelectedMovie(null);
    }

    return (
        <>
            <div className="movies-container">
                <div className="movies-search-container">
                    <p>Search your movie list</p>
                    <br />
                    <div className="movies-search-input-container">
                        <input className="movies-search-input-style" placeholder="Enter a keyword..." value={searchKeyword} onChange={(e) => setSearchKeyword(e.target.value)} />
                        <div className="movies-search-icon-container">
                            <img src={RightArrowIcon} alt="Right Arrow Icon" className="right-arrow-icon" onClick={() => filterMovies()}/>
                        </div>
                    </div>
                </div>
                <br />
                <br />
                <br />
                <div className="movie-list-header-container">
                    <div className="movie-list-header-left">
                        <p>My Movie List</p>
                        <button className="sign-in-button" onClick={() => setShowAddMovieModal(true)}>+ Add Movie</button>
                    </div>
                    <div className="movie-list-header-right">
                    </div>
                </div>
                <p style={{ fontSize: 14 }}><i>Check the movie if you have watched it already</i></p>
                <br />
                {movies.length <= 0 ? <div className="no-data-container">No Movies Added</div>
                : <div style={{ maxHeight: "calc(100vh - 370px)", overflowY: "auto" }}>
                    <table>
                        <tr>
                            <th style={{ width: "5%" }}></th>
                            <th style={{ width: "45%" }}>Movie Title</th>
                            <th style={{ width: "20%" }}>Date Added</th>
                            <th style={{ width: "20%" }}>Last Modified</th>
                            <th style={{ width: "10%" }}></th>
                        </tr>
                        {filteredMovies.map(movie => {
                            return <tr>
                                <td style={{ textAlign: "center" }}>
                                    <input type="checkbox" checked={movie.IsWatched} onChange={() => editMovie(movie)} />&nbsp;&nbsp;
                                </td>
                                <td>{movie.Title}</td>
                                <td>{movie.DateAdded}</td>
                                <td>{movie.LastModified}</td>
                                <td style={{ textAlign: "center" }}>
                                    <img 
                                        src={EditIcon} 
                                        alt="Edit Icon" 
                                        className="table-icon" 
                                        title="Edit" 
                                        onClick={() => {
                                            setShowAddMovieModal(true)
                                            setIsEditMovie(true);
                                            setTitle(movie.Title);
                                            setIsWatched(movie.IsWatched);
                                            setSelectedMovie(movie);
                                        }}
                                    />
                                    <img 
                                        src={DeleteIcon} 
                                        alt="Delete Icon" 
                                        className="table-icon" 
                                        title="Delete" 
                                        onClick={() => deleteMovie(movie.Id)}
                                    />
                                </td>
                            </tr>
                        })}
                    </table>
                </div>}
            </div>

            {showAddMovieModal &&
            <div className="add-movie-modal-layout">
                <div className="add-movie-modal-container">
                    <p><b>{isEditMovie ? "Edit" : "Add"} Movie</b></p>
                    <br />
                    <br />
                    <p className="input-header-style">Title</p>
                    <input className="add-movie-input-style" placeholder="Enter title" value={title} onChange={(e) => setTitle(e.target.value)} />
                    <br />
                    <br />
                    <p className="input-header-style">Have you watched it?</p>
                    <input type="radio" name="watched" value={"yes"} checked={isWatched === true} onChange={(e) => setIsWatched(e.target.value === "yes")} />&nbsp;&nbsp;<span style={{ fontSize: 14 }}>Yes</span>&nbsp;&nbsp;&nbsp;&nbsp;
                    <input type="radio" name="watched" value={"no"} checked={isWatched === false} onChange={(e) => setIsWatched(e.target.value === "yes")} />&nbsp;&nbsp;<span style={{ fontSize: 14 }}>No</span>
                    <br />
                    <br />
                    <br />
                    <button 
                        className="cancel-button left" 
                        onClick={() => resetAddMovieForm()}
                    >Cancel</button>
                    <button 
                        className="sign-in-button right" 
                        onClick={() => {
                            isEditMovie ? editMovie(selectedMovie) : addMovie()
                        }}
                    >{isEditMovie ? "Edit" : "Add"} Movie</button>
                </div>
            </div>}
        </>
    )
}

export default Movies;