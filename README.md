# Getting started with alinker

This project is for showcasing my Golang skills for building apis
Its live here [alinker] (<https://alinker.tk>)

alinker is url shortener similar to bit.ly for developers, It uses
redis for url mapping storage and mongodb for customers.

## Available endpoints

### `/docs`

for clear details and documentation of alinker api, how a developer can use it and develop using it.

### `/create-short-url`

Creates the shortened url and returns it, Expects a `long_url` key in a json request.
The user has to be Authenticated to access this endpoint, Authenticated user has a key `Authentication` having the returned key obtained in the header response after login.

### `/:shorturl`

Redirects the user to the long Url mapped with the given shorturl

### `/register`

Creates a new user in the database, The endpoint expects `username`, `email` and `password` in a json request

### `/login`

Expects `email`, `password` in a josn request.
Response has a JWT Authentication token in the header if a request is a success. The key is `Authentication`.

## TODOS

Some implemetation are still in progress

- logging
- some updates
