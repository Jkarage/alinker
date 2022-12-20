# Getting started with alinker

This project is for showcasing my Golang skills for building apis
Its live here [alinker] (<https://alinker.tk>)

alinker is url shortener similar to bit.ly for developers, It uses
redis for url mapping and mongodb for storing customers.

## Available endpoints

### `/register`

Creates a new user in the database, The endpoint expects `username`, `email` and `password` in a json formatted request body

[register.webm](https://user-images.githubusercontent.com/52887226/208636889-375674ce-897c-4c7f-9b25-14e27ec91510.webm)


### `/login`

The endpoint expects `email`, `password` in a json formatted request body.
Response has a JWT Authentication token in the header if a request is a success. The key is named `Authentication`.You are supposed to add it in all secure endpoints in the header with the given key `Authentication`.

[login.webm](https://user-images.githubusercontent.com/52887226/208636984-61c93f04-f8c0-43bc-9d35-988b97845cc8.webm)


### `/docs`

For clear details and documentation of alinker api. The docs page is still in progress.

### `/create-short-url`

Creates the shortened url and returns it, Expects a `long_url` key in a json request.
The user has to be Authenticated to access this endpoint, Authenticated user has a key `Authentication` having the returned key obtained in the header response after login.

[create-short-url.webm](https://user-images.githubusercontent.com/52887226/208637077-10012ea2-1e9e-44c7-bb09-c2aef5d18a05.webm)

### `/:shorturl`

Redirects the user to the long Url mapped with the given shorturl, Looked up on the redis database.

## TODOS

Some implemetation are still in progress

- logging
- some updates
