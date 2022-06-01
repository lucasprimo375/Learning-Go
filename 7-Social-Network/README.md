# How to set up The Social Network

## Pre-Requisites
- MySQL

## How to configure
1. Create database and a user on MySQL with a password and grant it all privileges in your database.
2. Create .env file inside the api folder and add the following environment variables with the values created in earlier steps:
    - API_PORT
    - DB_USER
    - DB_PASSWORD
    - DB_NAME
3. Create .env file inside the webapp folder and add the following environment variables with the values created in earlier steps:
    - API_URL
    - APP_PORT
    - HASH_KEY
    - BLOCK_KEY