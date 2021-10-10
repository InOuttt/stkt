package movie

import (
	"testing"

	"github.com/inouttt/stkt/test_2/internal/config"
	"github.com/inouttt/stkt/test_2/internal/log"
)

func TestSearch(t *testing.T) {
	dbDsn := "root:pass@/test"
	urlTarget := "http://www.omdbapi.com/"

	args := config.Args{
		MysqlDsn:  dbDsn,
		DbTimeout: 10,
	}

	tableTest := []struct {
		name           string
		req            SearchRequest
		expectedResult string
		expectedError  error
	}{
		{
			name: "TEST SUCCESS SEARCH",
			req: SearchRequest{
				Apikey: "faf7e5bb",
				Search: "batman",
				Page:   "2",
			},
			expectedResult: `{"Search":[{"Title":"Batman: The Killing Joke","Year":"2016","imdbID":"tt4853102","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"Title":"Batman: The Dark Knight Returns, Part 2","Year":"2013","imdbID":"tt2166834","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BYTEzMmE0ZDYtYWNmYi00ZWM4LWJjOTUtYTE0ZmQyYWM3ZjA0XkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_SX300.jpg"},{"Title":"Batman: Mask of the Phantasm","Year":"1993","imdbID":"tt0106364","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BYTRiMWM3MGItNjAxZC00M2E3LThhODgtM2QwOGNmZGU4OWZhXkEyXkFqcGdeQXVyNjExODE1MDc@._V1_SX300.jpg"},{"Title":"Batman: Year One","Year":"2011","imdbID":"tt1672723","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNTJjMmVkZjctNjNjMS00ZmI2LTlmYWEtOWNiYmQxYjY0YWVhXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"Title":"Batman: Assault on Arkham","Year":"2014","imdbID":"tt3139086","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BZDU1ZGRiY2YtYmZjMi00ZDQwLWJjMWMtNzUwNDMwYjQ4ZTVhXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"Title":"Batman: The Movie","Year":"1966","imdbID":"tt0060153","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMmM1OGIzM2UtNThhZS00ZGNlLWI4NzEtZjlhOTNhNmYxZGQ0XkEyXkFqcGdeQXVyNTkxMzEwMzU@._V1_SX300.jpg"},{"Title":"Batman: Gotham Knight","Year":"2008","imdbID":"tt1117563","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BM2I0YTFjOTUtMWYzNC00ZTgyLTk2NWEtMmE3N2VlYjEwN2JlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"Title":"Batman: Arkham City","Year":"2011","imdbID":"tt1568322","Type":"game","Poster":"https://m.media-amazon.com/images/M/MV5BZDE2ZDFhMDAtMDAzZC00ZmY3LThlMTItMGFjMzRlYzExOGE1XkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"Title":"Batman Beyond","Year":"1999–2001","imdbID":"tt0147746","Type":"series","Poster":"https://m.media-amazon.com/images/M/MV5BYTBiZjFlZDQtZjc1MS00YzllLWE5ZTQtMmM5OTkyNjZjMWI3XkEyXkFqcGdeQXVyMTA1OTEwNjE@._V1_SX300.jpg"},{"Title":"Son of Batman","Year":"2014","imdbID":"tt3139072","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BYjdkZWFhNzctYmNhNy00NGM5LTg0Y2YtZWM4NmU2MWQ3ODVkXkEyXkFqcGdeQXVyNTA0OTU0OTQ@._V1_SX300.jpg"}],"totalResults":"463","Response":"True"}`,
			expectedError:  nil,
		},
		{
			name: "TEST WRONG KEY",
			req: SearchRequest{
				Apikey: "wrongkey",
				Search: "batman",
				Page:   "2",
			},
			expectedResult: `{"Response":"False","Error":"Invalid API key!"}`,
			expectedError:  nil,
		},
	}

	dbCfg := config.GetDbConf(args)
	mysqlS, _ := config.NewMysqlSession(dbCfg)
	lr := log.NewRepository(mysqlS)
	ls := log.Newservices(lr)
	ms := NewServices(urlTarget, ls)

	for _, test := range tableTest {
		resp, err := ms.Search(test.req)

		if err != test.expectedError {
			t.Error(test.name + " Invalid Error")
		}

		if resp != test.expectedResult {
			t.Log(resp)
			t.Error(test.name + " Invalid result")
		}
	}

}

func TestDetail(t *testing.T) {
	dbDsn := "root:pass@/test"
	urlTarget := "http://www.omdbapi.com/"

	args := config.Args{
		MysqlDsn:  dbDsn,
		DbTimeout: 10,
	}

	tableTest := []struct {
		name           string
		req            DetailRequest
		expectedResult string
		expectedError  error
	}{
		{
			name: "TEST SUCCESS SEARCH",
			req: DetailRequest{
				Apikey: "faf7e5bb",
				Id:     "tt4853102",
			},
			expectedResult: `{"Title":"Batman: The Killing Joke","Year":"2016","Rated":"R","Released":"25 Jul 2016","Runtime":"76 min","Genre":"Animation, Action, Crime","Director":"Sam Liu","Writer":"Brian Azzarello, Brian Bolland, Bob Kane","Actors":"Kevin Conroy, Mark Hamill, Tara Strong","Plot":"As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.","Language":"English","Country":"United States","Awards":"1 win & 2 nominations","Poster":"https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"6.4/10"},{"Source":"Rotten Tomatoes","Value":"39%"}],"Metascore":"N/A","imdbRating":"6.4","imdbVotes":"52,878","imdbID":"tt4853102","Type":"movie","DVD":"02 Aug 2016","BoxOffice":"$3,775,000","Production":"Warner Bros.","Website":"N/A","Response":"True"}`,
			expectedError:  nil,
		},
		{
			name: "TEST WRONG Id",
			req: DetailRequest{
				Apikey: "faf7e5bb",
				Id:     "wrongid",
			},
			expectedResult: `{"Response":"False","Error":"Incorrect IMDb ID."}`,
			expectedError:  nil,
		},
		{
			name: "TEST empty id",
			req: DetailRequest{
				Apikey: "faf7e5bb",
			},
			expectedResult: `{"Response":"False","Error":"Incorrect IMDb ID."}`,
			expectedError:  nil,
		},
		{
			name: "TEST No api key",
			req: DetailRequest{
				Id: "wrongid",
			},
			expectedResult: `{"Response":"False","Error":"No API key provided."}`,
			expectedError:  nil,
		},
	}

	dbCfg := config.GetDbConf(args)
	mysqlS, _ := config.NewMysqlSession(dbCfg)
	lr := log.NewRepository(mysqlS)
	ls := log.Newservices(lr)
	ms := NewServices(urlTarget, ls)

	for _, test := range tableTest {
		resp, err := ms.Detail(test.req)

		if err != test.expectedError {
			t.Error(test.name + " Invalid Error")
		}

		if resp != test.expectedResult {
			t.Log(resp)
			t.Error(test.name + " Invalid result")
		}
	}

}
