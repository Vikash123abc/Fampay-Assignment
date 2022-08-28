# Fampay-Assignment

### Make an API to fetch latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

# Language : Golang
# DataBase: MongoDB
  ( Having index on title and description field )


# APIS

1.  Returns list of videos with pagination. ( GET Method )

    >  http://localhost:2000/?offset=1&limit=3

2.  List of videos with regex matching on searchString with pagination . (GET Method)
   > Eg: http://localhost:2000/search?limit=4&offset=2&search=cricket

### Response Format

`{ "success": boolean, "videos": []string, "error": string }`

# Functionalities

- There is a `go routine` which gets videos from youtube every 10 seconds with predefined search string ("cricket" in this example).

- User can supply `multiple API keys`, first valid API key in the list will be used everytime a request is made.

- searchString does the regular expression matching on title or description. The search is case insensitive.



# Config file

Update the config file like belowing:

MONGO_URI: 

YOUTUBE_API_KEY: 


# Running the server

1. `go run main.go` will start the server locally on port 2000.
   or
2. `docker build . -t main && docker run -p 2000:2000 main` will start a docker container on port 2000.
