import './App.css';
import {useEffect, useState} from "react";

function App() {
    const [books, setBooks] = useState([]);
    const [libraries, setLibraries] = useState([]);
    const [publishers, setPublishers] = useState([]);
    const [users, setUsers] = useState([]); 
    let title, genre, review = '';

    useEffect(() => {
        (async () => {
            const response = await fetch('http://localhost:3005/api/books');

            const content = await response.json()

            setBooks(content);
        })()
    }, []);

    useEffect(() => {
        (async () => {
            const response = await fetch('http://localhost:3015/api/libraries');

            const contentlib = await response.json()

            setLibraries(contentlib);
        })()
    }, []);

    useEffect(() => {
        (async () => {
            const response = await fetch('http://localhost:3020/api/publishers');

            const contentpub = await response.json()

            setPublishers(contentpub);
        })()
    }, []);

    useEffect(() => {
        (async () => {
            const response = await fetch('http://localhost:3030/api/users');

            const contentuser = await response.json()

            setUsers(contentuser);
        })()
    }, []);

    const createBook = async e => {
        e.preventDefault();

        const res = await fetch('http://localhost:3005/api/books', {
            method: "POST",
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                title,
                genre,
            })
        });

        const createdBook = await res.json();

        setBooks([...books, createdBook]);
    }

    const createReview = async (e, book_id) => {
        e.preventDefault();

        const response = await fetch('http://localhost:3010/api/reviews', {
            method: "POST",
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                book_id,
                text: review
            })
        });

        const createdReview = await response.json();

        review = ''

        setBooks(books.map(p => {
            if (p.id === book_id) {
                p.reviews.push(createdReview)
            }

            return p;
        }))
    }
    return (
        
        <div className="App container">
            <nav class="navbar navbar-expand-lg navbar-light bg-light">
            <div class="container-fluid">
                <a class="navbar-brand" href="#nav">BookReviews</a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarColor02" aria-controls="navbarColor02" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
                </button>

                <div class="collapse navbar-collapse" id="navbarColor02">
                <ul class="navbar-nav me-auto">
                    <li class="nav-item">
                    <a class="nav-link active" href="#home">Home
                        <span class="visually-hidden">(current)</span>
                    </a>
                    </li>
                    <li class="nav-item">
                    <a class="nav-link" href="#about">About</a>
                    </li>
                </ul>
                </div>
                <div>
                {users.length && users.map(
                    user => {
                        return (
                            <div className="py-2">
                                <h6 className="my-0 fw-normal">
                                    <a href="#filterby"><svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" fill="currentColor" class="bi bi-person-circle mx-2 mb-1" viewBox="0 0 16 16">
                                    <path d="M11 6a3 3 0 1 1-6 0 3 3 0 0 1 6 0z"/>
                                    <path fill-rule="evenodd" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8zm8-7a7 7 0 0 0-5.468 11.37C3.242 11.226 4.805 10 8 10s4.757 1.225 5.468 2.37A7 7 0 0 0 8 1z"/>
                                    </svg></a>
                                    {user.username}
                                </h6>
                            </div>
                        )
                    }
                )}
                </div>
            </div>
            </nav>
            <div className="row my-5">
                <div className="col-2">
                    <div className="col">
                        <div className="card mb-3 rounded-3 shadow-sm">
                            <div className="card-header py-3">
                                <h4 className="my-0 fw-normal">Libraries</h4>
                            </div>
                            {libraries.length && libraries.map(
                                library => {
                                    return (
                                        <div className="card-footer py-2">
                                            <h6 className="my-0 fw-normal">
                                                {library.name}, {library.location}
                                                <a href="#filterby"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-funnel float-end" viewBox="0 0 16 16">
                                                    <path d="M1.5 1.5A.5.5 0 0 1 2 1h12a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.128.334L10 8.692V13.5a.5.5 0 0 1-.342.474l-3 1A.5.5 0 0 1 6 14.5V8.692L1.628 3.834A.5.5 0 0 1 1.5 3.5v-2zm1 .5v1.308l4.372 4.858A.5.5 0 0 1 7 8.5v5.306l2-.666V8.5a.5.5 0 0 1 .128-.334L13.5 3.308V2h-11z"/>
                                                </svg></a>
                                            </h6>
                                        </div>
                                    )
                                }
                            )}
                        </div>
                    </div>
                    <div className="col">
                        <div className="card mb-3 rounded-3 shadow-sm">
                            <div className="card-header py-3">
                                <h4 className="my-0 fw-normal">Publishers</h4>
                            </div>
                            {publishers.length && publishers.map(
                                publisher => {
                                    return (
                                        <div className="card-footer py-2">
                                            <h6 className="my-0 fw-normal">
                                                {publisher.name}
                                                <a href="#filterby"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-funnel float-end" viewBox="0 0 16 16">
                                                    <path d="M1.5 1.5A.5.5 0 0 1 2 1h12a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.128.334L10 8.692V13.5a.5.5 0 0 1-.342.474l-3 1A.5.5 0 0 1 6 14.5V8.692L1.628 3.834A.5.5 0 0 1 1.5 3.5v-2zm1 .5v1.308l4.372 4.858A.5.5 0 0 1 7 8.5v5.306l2-.666V8.5a.5.5 0 0 1 .128-.334L13.5 3.308V2h-11z"/>
                                                </svg></a>
                                            </h6>
                                        </div>
                                    )
                                }
                            )}
                        </div>
                    </div>
                </div>
            <div className='col-10'>
            <form className="row" onSubmit={createBook}>
                <div className="col-4">
                    <h2>Create a Book</h2>

                    <input className="form-control mb-3" onChange={e => title = e.target.value}/>
                    <input className="form-control mb-3" onChange={e => genre = e.target.value}/>

                    <button className="btn btn-primary" type="submit">Save</button>
                </div>
            </form>

            <main>
                <div className="row row-cols-1 row-cols-md-3 mb-3 text-center my-5">
                    {books.length && books.map(
                        book => {
                            return (
                                <div className="col" key={book.id}>
                                    <div className="card mb-4 rounded-3 shadow-sm">
                                        <div className="card-header py-3">
                                            <h4 className="my-0 fw-normal">{book.title}</h4>
                                        </div>
                                        <div className="card-body">
                                            <p className="card-title pricing-card-title">{book.genre}</p>
                                            <form onSubmit={e => createReview(e, book.id)}>
                                                <input className="w-100 form-control"
                                                       onChange={e => review = e.target.value}/>
                                            </form>
                                        </div>
                                        {book.reviews && book.reviews.map(
                                            review => {
                                                return (
                                                    <div className="card-footer py-3" key={review.id}>
                                                        {review.text}
                                                    </div>
                                                )
                                            }
                                        )}
                                    </div>
                                </div>
                            )
                        }
                    )}
                </div>
            </main>
        </div>
            </div>
            </div>
            
            
    );
}

export default App;
