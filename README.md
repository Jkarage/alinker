# Getting started with alinker

This project is for showcasing my Golang skills for building apis
Its live here [alinker] (<https://alinker.tk>)

alinker is url shortener similar to bit.ly for developers, It uses
redis for url mapping storage and mongodb for customers.

## Available endpoints

### `/docs`

for clear details and documentation of alinker api, how a developer can use it and develop using it.

### `/create-short-url`

Creates the shortened url and returns it, Expects a `long_url` and `user_id` keys in a json request.

### `/:shorturl`

Redirects the user to the long Url mapped with the given shorturl

### `/register`

Creates a new user in the database, Read the [docs] (<https://alinker.tk/docs>) for more on the signature of the request.

### `/login`

Returns an JWT Authentication token in the header if a request is a success.

## TODOS

Some implemetation are still in progress

- logging
- some updates
