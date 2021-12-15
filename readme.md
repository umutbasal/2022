# 2022

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

[![asciicast](https://asciinema.org/a/htM9lWdlMREFeDIq724Qy5OAA.svg)](https://asciinema.org/a/htM9lWdlMREFeDIq724Qy5OAA)

## Running
`go run .`

## Testing
`curl http://localhost:8080`

## Docker with ko (mocked with my iterm specs)
`docker run --rm -p 8080:8080 $(ko publish . --local)`
