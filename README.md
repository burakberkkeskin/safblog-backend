## Overview

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
# JWT Values
JWTSECRET="secretvalue"
JWTHOUR="72"
# Root User Values
ROOT_USER_USERNAME="safderun"
ROOT_USER_EMAIL="user@gmail.com"
ROOT_USER_PASSWORD="Test1234"
# Database Values
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

## Request Examples

- You can get and import the request examples from `./thunder-requests` directory.

- Install the Thunder Client.

  - Extension id: `rangav.vscode-thunder-client`

- Import the collection from the directory.

## Docker image.

### Build

```bash
docker image build -t safblog-backend:latest .
```

### Run

```bash
docker run --name safblog-backend --restart always -d --env-file .env --network safblog -p 8080:8080 safblog-backend:latest
```

## Routes

You can find the routes under the directory.

@TODO
