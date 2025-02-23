# Codeforces RSS Feed Generator

## Overview
`codeforces-rss` is a lightweight Go-based service that generates an RSS feed with random Codeforces problems based on specified **tags** and **difficulty ratings**. The feed can be consumed by any RSS reader, such as [FreshRSS](https://freshrss.org/), allowing you to receive daily problem recommendations for consistent practice.

## Features
✅ Fetches **random** Codeforces problems based on given tags and rating range.  
✅ Exposes an **RSS feed endpoint** (`/rss`).  
✅ Customizable **tags, minimum & maximum rating** via environment variables.  
✅ Lightweight and **Docker-ready** for easy deployment.  

## Installation

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/yourusername/codeforces-rss.git
cd codeforces-rss
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Run Locally
```sh
export CF_TAGS="dp,graphs"
export CF_MIN_RATING="800"
export CF_MAX_RATING="1500"
go run cmd/rss-server/main.go
```

### 4️⃣ Test RSS Feed
```sh
curl http://localhost:8080/rss
```

## Configuration
The service uses **environment variables** for configuration:

| Environment Variable  | Description | Default |
|----------------------|-------------|---------|
| `CF_TAGS`           | Comma-separated list of Codeforces tags (e.g., `dp,greedy`) | `dp,graphs` |
| `CF_MIN_RATING`     | Minimum problem rating | `800` |
| `CF_MAX_RATING`     | Maximum problem rating | `1500` |

Alternatively, a fallback configuration file exists at `config/config.json`.

## Docker Deployment

### 1️⃣ Build Docker Image
```sh
docker build -t codeforces-rss .
```

### 2️⃣ Run the Container
```sh
docker run -d --name cf-rss \
  -e CF_TAGS="dp,greedy" \
  -e CF_MIN_RATING="900" \
  -e CF_MAX_RATING="1600" \
  -p 8080:8080 \
  codeforces-rss
```

### 3️⃣ Test RSS Feed
```sh
curl http://localhost:8080/rss
```

## Docker Compose (Optional)

Create a `docker-compose.yml` file:

```yaml
version: '3.8'
services:
  rss-server:
    build: .
    container_name: cf-rss
    environment:
      - CF_TAGS=dp,greedy
      - CF_MIN_RATING=900
      - CF_MAX_RATING=1600
    ports:
      - "8080:8080"
    restart: unless-stopped
```

Run it with:
```sh
docker-compose up -d
```

