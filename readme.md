# 2022

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

[![CodeQL](https://github.com/umutbasal/2022/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/umutbasal/2022/actions/workflows/codeql-analysis.yml)

[![asciicast](https://asciinema.org/a/tp64MGwT8WtkHaNPOM12mJGGc.svg)](https://asciinema.org/a/tp64MGwT8WtkHaNPOM12mJGGc)

## Running
`go run .`

## Testing
`curl http://localhost:8080`

## Docker with ko (mocked with my iterm specs)
`docker run --rm -p 8080:8080 $(ko publish . --local)`
