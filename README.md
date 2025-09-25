## 📝 URL Shortener API

A simple URL shortening service built with Go and PostgreSQL.
It allows users to create short URLs for long URLs and redirect users from the short URL to the original URL.

---

## 📦 Features

Shorten long URLs to a 6-character code.

Redirect from short URL to original URL.

Validate for duplicate URLs and short codes.

Uses PostgreSQL to store URL mappings.

RESTful API endpoints.

---

## ⚙️ Prerequisites

Go 1.20+

PostgreSQL 15+

github.com/lib/pq PostgreSQL driver

---

## 🛠️ Setup Instructions

### 1. Clone the repository
    git clone https://github.com/Aashritha123-lab/GO_BASICS.git
    cd shorten_url
---
### 2. Create PostgreSQL database & table
    CREATE DATABASE url_shortener;

    \c url_shortener

    CREATE TABLE short_url (
        id SERIAL PRIMARY KEY,
        short_response VARCHAR(6) UNIQUE NOT NULL,
        url TEXT UNIQUE NOT NULL
    );
---

### 3. Update DB connection

    Edit the DBConnection() function in main.go:

    dsn := "user=username password=example@123 dbname=dbname sslmode=disable"


Replace with your PostgreSQL credentials.
---
### 4. Run the server
    
    go run main.go


    Server will start on http://localhost:3051.

## 🔗 API Endpoints

### 1. Shorten URL

    URL: /shorten

    Method: POST

    Body (JSON):

    {
        "url": "https://example.com"
    }


    Response (JSON):

    {
        "short_url": "http://localhost:3051/Ab12Cd",
        "original_url": {
            "url": "https://example.com"
        }
    }

### 2. Redirect Short URL

    URL: /{short_code}

    Method: GET

    Redirects the user to the original URL.
---

## 🔄 URL Shortener Flow Diagram

flowchart TD
    A[Client: POST /shorten] -->|Send original URL| B[Go Server: ShorthandUrl Handler]
    B --> C{Validate URL}
    C -->|URL exists| D[Return 409 Conflict]
    C -->|URL does not exist| E[Generate short code]
    E --> F{Check code in DB}
    F -->|Code exists| E
    F -->|Code unique| G[Insert short_url into DB]
    G --> H[Return JSON with short URL]
    
    I[Client: GET /{short_code}] -->|Request short URL| J[Go Server: Redirect Handler]
    J --> K[Check short_code in DB]
    K -->|Found| L[HTTP Redirect to original URL]
    K -->|Not found| M[Return 404 Not Found]

## ⚠️ Notes

    Short codes are 6-character alphanumeric strings.

    Duplicate long URLs will return a conflict error.

    Random short codes are generated; collisions are checked in the database.

## 🛠️ Improvements / Next Steps

    Add expiration date for short URLs.

    Track analytics (click counts, timestamps).

    Add user authentication.

    Dockerize the app for easy deployment.