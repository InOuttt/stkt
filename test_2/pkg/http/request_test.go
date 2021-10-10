package http

import (
	"testing"
)

func TestGetRequest(t *testing.T) {
	tableTest := []struct {
		name           string
		url            string
		qry            map[string]string
		expectedResult string
		expectedError  error
	}{
		{
			name: "TEST SUCCESS SEARCH",
			url:  "http://www.omdbapi.com/",
			qry: map[string]string{
				"apikey": "faf7e5bb",
				"s":      "batman",
				"p":      "2",
			},
			expectedResult: `{"Search":[{"Title":"Batman Begins","Year":"2005","imdbID":"tt0372784","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"},{"Title":"Batman v Superman: Dawn of Justice","Year":"2016","imdbID":"tt2975590","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"},{"Title":"Batman","Year":"1989","imdbID":"tt0096895","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTYwNjAyODIyMF5BMl5BanBnXkFtZTYwNDMwMDk2._V1_SX300.jpg"},{"Title":"Batman Returns","Year":"1992","imdbID":"tt0103776","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BOGZmYzVkMmItM2NiOS00MDI3LWI4ZWQtMTg0YWZkODRkMmViXkEyXkFqcGdeQXVyODY0NzcxNw@@._V1_SX300.jpg"},{"Title":"Batman Forever","Year":"1995","imdbID":"tt0112462","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNDdjYmFiYWEtYzBhZS00YTZkLWFlODgtY2I5MDE0NzZmMDljXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg"},{"Title":"Batman & Robin","Year":"1997","imdbID":"tt0118688","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMGQ5YTM1NmMtYmIxYy00N2VmLWJhZTYtN2EwYTY3MWFhOTczXkEyXkFqcGdeQXVyNTA2NTI0MTY@._V1_SX300.jpg"},{"Title":"The Lego Batman Movie","Year":"2017","imdbID":"tt4116284","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTcyNTEyOTY0M15BMl5BanBnXkFtZTgwOTAyNzU3MDI@._V1_SX300.jpg"},{"Title":"Batman: The Animated Series","Year":"1992–1995","imdbID":"tt0103359","Type":"series","Poster":"https://m.media-amazon.com/images/M/MV5BOTM3MTRkZjQtYjBkMy00YWE1LTkxOTQtNDQyNGY0YjYzNzAzXkEyXkFqcGdeQXVyOTgwMzk1MTA@._V1_SX300.jpg"},{"Title":"Batman: Under the Red Hood","Year":"2010","imdbID":"tt1569923","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNmY4ZDZjY2UtOWFiYy00MjhjLThmMjctOTQ2NjYxZGRjYmNlL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"Title":"Batman: The Dark Knight Returns, Part 1","Year":"2012","imdbID":"tt2313197","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMzIxMDkxNDM2M15BMl5BanBnXkFtZTcwMDA5ODY1OQ@@._V1_SX300.jpg"}],"totalResults":"463","Response":"True"}`,
			expectedError:  nil,
		},
		{
			name: "TEST SUCCESS DETAIL",
			url:  "http://www.omdbapi.com/",
			qry: map[string]string{
				"apikey": "faf7e5bb",
				"i":      "tt4853102",
			},
			expectedResult: `{"Title":"Batman: The Killing Joke","Year":"2016","Rated":"R","Released":"25 Jul 2016","Runtime":"76 min","Genre":"Animation, Action, Crime","Director":"Sam Liu","Writer":"Brian Azzarello, Brian Bolland, Bob Kane","Actors":"Kevin Conroy, Mark Hamill, Tara Strong","Plot":"As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.","Language":"English","Country":"United States","Awards":"1 win & 2 nominations","Poster":"https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"6.4/10"},{"Source":"Rotten Tomatoes","Value":"39%"}],"Metascore":"N/A","imdbRating":"6.4","imdbVotes":"52,878","imdbID":"tt4853102","Type":"movie","DVD":"02 Aug 2016","BoxOffice":"$3,775,000","Production":"Warner Bros.","Website":"N/A","Response":"True"}`,
			expectedError:  nil,
		},
		{
			name: "TEST WRONG KEY",
			url:  "http://www.omdbapi.com/",
			qry: map[string]string{
				"apikey": "wrongKey",
				"i":      "tt4853102",
			},
			expectedResult: `{"Response":"False","Error":"Invalid API key!"}`,
			expectedError:  nil,
		},
	}

	for _, test := range tableTest {
		resp, err := Get(test.url, test.qry)

		if err != test.expectedError {
			t.Error(test.name + " Invalid Error")
		}

		if resp != test.expectedResult {
			t.Error(test.name + " Invalid result")
		}
	}

}
