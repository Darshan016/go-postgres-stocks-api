# Go Stocks API

A simple API for managing stock information using Go and PostgreSQL.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)

## Introduction

The Go Stocks API is a web application that allows users to manage stock information. The API is built using Go and connects to a PostgreSQL database to store stock data.

## Features

- Create new stock entries
- Retrieve stock details
- Update stock information
- Delete stock entries

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Darshan016/go-postgres-stocks-api.git
    cd go-postgres-stocks-api
    ```

2. Install Go dependencies:
    ```sh
    go mod tidy
    ```

3. Set up PostgreSQL and create a database:
    ```sql
    CREATE DATABASE stocks_db;
    ```

4. Create the necessary table in the database:
    ```sql
    CREATE TABLE stocks (
        stockid SERIAL PRIMARY KEY,
        name VARCHAR(100),
        price NUMERIC,
        company VARCHAR(100)
    );
    ```

## Configuration

Create a `.env` file in the project root and add the following configuration:
```sh
POSTGRES_URL=your_postgres_url
