## Overview

Safblog backend service.

## Run the app (Development)

- Clone the repository.

```bash
git clone https://github.com/safderun/safblog-backend.git && \
cd safblog-backend
```

- Install `air` tool.

```bash
go install github.com/cosmtrek/air@latest
```

- Run the dependencies from the `docker-compose.yaml`

```bash
docker compose up -d
```

- Create a `.env` file.

```env
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=postgresuser
DB_PASSWORD=postgrespasswd
DB_NAME=safblog
```

- Run `air` in the repo

```bash
air
```

## Docker image.

### Build

```bash
docker image build -t safblog-backend:latest .
```

### Run

```bash
docker run --name safblog-backend --restart always safblog-backend:latest
```

## Routes

You can find the routes under the directory.

@TODO
