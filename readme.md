# 2022

[![asciicast](https://asciinema.org/a/bmfyVBgOZQDN8tKIOMzc8QwYV.svg)](https://asciinema.org/a/bmfyVBgOZQDN8tKIOMzc8QwYV)

## Running
`go run .`

## Testing
`curl http://localhost:8080`

## Docker with ko (mocked with my iterm specs)
`docker run --rm -p 8080:8080 $(ko publish . --local)`
