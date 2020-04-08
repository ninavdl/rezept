A virtual recipe book, written in Go and Vue.js

# Running rezept

## With docker-compose (preferred)

The easiest way to setup rezept is by using docker-compose.

1. Create a `docker-compose.yml` file:
```yaml
version: '3'
services:
  frontend:
    image: 'rezept-frontend:${VERSION}'
    environment:
      - API_URL=/api
      - BASE_URL=/
      - PAGE_TITLE=${PAGE_TITLE}
  backend:
    image: 'rezept-backend:${VERSION}'
    environment:
      - PATH_PREFIX=/
      - API_PREFIX=/
      - BIND_ADDRESS=0.0.0.0:8080
      - DB_PATH=/data/rezept.sqlite
      - IMAGE_PATH=/data/images
      - IMAGE_URL=/api/images
    volumes:
      - "./data:/data"
  reverse-proxy:
    image: 'rezept-reverse-proxy:${VERSION}'
    ports:
      - "${LISTEN}:80"
    depends_on:
      - frontend
      - backend
```

2. Create a `.env` file (you can change the parameters to your likening):
```
PAGE_TITLE=rezept
LISTEN=8080
VERSION=latest
```
3. Run with `docker-compose up` - the application should be available on port 8080.

## Running with docker (manual)

### Backend
The backend docker container is available as `rezept/rezept-backend`. You should specify the following environment variables:
- `BIND_ADDRESS`: On which address to bind the HTTP server, `host:port` format
- `DB_PATH`: Where the database should be stored - ideally in a docker volume
- `IMAGE_PATH`: Where images should be stored - ideally in a docker volume
- `IMAGE_URL`: Prefix which should be prepended to image urls - it should correspond to where the /images resource will be available

### Frontend
The frontend docker container is available as `rezept-frontend`. You should specify the following environment variables:
- `API_URL`: Base url of the API (rezept-backend), defaults to `/api`.
- `BASE_URL`: Must be `/` for now. If you want to serve the frontend under a different base path `/$PREFIX`, you have to adapt the Dockerfile and change `RUN npm run build` to `RUN npx parcel build --no-source-maps --no-autoinstall --public-url '/$PREFIX' src/index.html` and then adapt this parameter. Defaults to `/`.
- `PAGE_TITLE`: Name of this application, defaults to `rezept`.

## Running manually
Clone this repository and then build the backend and frontend separately:

### Backend
1. `cd backend`
2. `go mod download -x`
3. `go install -v ./...`
4. Create a `config.json` file:
```json
{
    "PathPrefix": "/",
    "APIPrefix": "/",
    "Address": "On which address to bind the HTTP server, host:port format",
    "DBPath": "Path to where the sqlite databse should be stored",
    "ImagePath": "Path to directory where images should be stored",
    "ImageURL": "Prefix which should be prepended to image urls - it should correspond to where the /images resource will be available"
}
```
5. Run `rezept [path-to-config.json]`

## Frontend

1. `cd frontend`
2. `npm install`
3. Modify `.env` file:
```
PAGE_TITLE="Name of this application"
BASE_URL="/"
API_URL="Path to where the API is reachable"
```
4. `npm run build`
5. Assets are available in the `dist` folder - deploy them on a web server of your preference.

(If you want to serve the frontend on a path different than `/`: Modify `BASE_URL` and run `npx parcel build --no-source-maps --no-autoinstall --public-url '/$BASE_URL' src/index.html` instead of `npm run build`)