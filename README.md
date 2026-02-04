# SkyMatch API

Backend service for astronomical image plate-solving. This API identifies celestial objects and constellations in
uploaded sky images using astrometry.net's Nova solver and SIMBAD database.

Frontend repository: [SkyMatch](https://github.com/oadultradeepfield/SkyMatch)

## Architecture

The project uses a two-layer architecture:

### Cloudflare Worker (`src/index.ts`)

Entry point that routes incoming requests to the container infrastructure.

### Go Container (`container_src/`)

The actual API server running on port 8080:

- **Router**: Chi with IP-based rate limiting (100 req/s)
- **External Services**: Nova (astrometry.net) for plate solving, SIMBAD for object identification
- **Structure**: MVC-like organization
    - `controller/` - HTTP handlers
    - `service/` - Business logic
    - `client/` - External API clients (Nova, SIMBAD)
    - `model/` - Data structures
    - `view/` - Response formatting

## Setup

```bash
pnpm install
```

## Local Development

For local testing, run the Go server directly:

```bash
cd container_src
go run ./cmd/server
```

Server runs on `localhost:8080`.

Note: `pnpm run dev` runs the Cloudflare Worker which requires container infrastructure.

## API Endpoints

| Method   | Endpoint              | Description                    |
|:---------|:----------------------|:-------------------------------|
| `GET`    | `/`                   | Health check                   |
| `GET`    | `/api/constellations` | Search constellations          |
| `POST`   | `/api/solve`          | Submit image for plate solving |
| `GET`    | `/api/solve/{jobId}`  | Get solve status               |
| `DELETE` | `/api/solve/{jobId}`  | Cancel solve job               |

## Deployment

```bash
pnpm run deploy
```

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).
