# openswingsonic

Open Subsonic API translation layer for [Swing Music](https://github.com/swingmx/swingmusic). Lets any OpenSubsonic-compatible client (Feishin, DSub, Symfonium, etc.) connect to a Swing Music server.

> **Note:** Swing Music does not expose a "get single track" endpoint in its stock API. Until that is upstreamed, openswingsonic requires the patched version from [swingmx/swingmusic#478](https://github.com/swingmx/swingmusic/pull/478).

## Quick start (Docker Compose)

```yaml
services:
  oswingsonic:
    image: ghcr.io/tikhonp/openswingsonic:latest
    container_name: oswingsonic
    environment:
      - SWINGMUSIC_BASE_URL=swingmusic:1970
      - PUBLIC_SWINGMUSIC_URL=https://music.example.com
      - CRED_PROVIDER=env
      - OSM_USER_0_USERNAME=alice
      - OSM_USER_0_PASSWORD=secret
    ports:
      - "1991:1991"
    volumes:
      - ./storage:/storage
```

Point your OpenSubsonic client at `http://<host>:1991` and log in with the credentials you configured.

## Configuration

| Variable | Required | Default | Description |
|---|---|---|---|
| `SWINGMUSIC_BASE_URL` | Yes | — | Internal URL of the Swing Music server |
| `DATABASE_PATH` | No | `/storage/openswingmusic.db` | Path to the SQLite state database |
| `PUBLIC_SWINGMUSIC_URL` | No | `SWINGMUSIC_BASE_URL` | Public URL of Swing Music, used in image URLs returned to clients |
| `CRED_PROVIDER` | No | `database` | Credentials provider: `database`, `file`, or `env` |
| `USERS_FILE_PATH` | If `CRED_PROVIDER=file` | — | Path to the users file |
| `LISTEN_ADDR` | No | `:1991` | TCP address to listen on |
| `DEBUG` | No | `false` | Enable verbose debug logging |
| `JSON_LOG` | No | `false` | Emit logs as JSON |

## Credentials providers

openswingsonic keeps its own user list separate from Swing Music's internal users. When a client logs in, the credentials are checked against this list and then used to open a session with Swing Music on the client's behalf.

### `env`

Users are read from environment variables at startup. Define as many users as needed by incrementing the index:

```
OSM_USER_0_USERNAME=alice
OSM_USER_0_PASSWORD=secret
OSM_USER_1_USERNAME=bob
OSM_USER_1_PASSWORD=hunter2
```

### `file`

Users are read from a plain-text file. Set `USERS_FILE_PATH` to the file's path. Each line is `username:password`:

```
alice:secret
bob:hunter2
```

### `database` (default)

Users are stored in the SQLite database. Add users directly in the database:

```sql
INSERT INTO swingmusic_users (username, password) VALUES ('alice', 'secret');
```

## Development

```yaml
services:
  oswingsonic:
    container_name: oswingsonic
    build:
      context: .
      target: dev
    environment:
      - DEBUG=true
      - SWINGMUSIC_BASE_URL=swingmusic:1970
      - PUBLIC_SWINGMUSIC_URL=https://music.example.com
      - CRED_PROVIDER=file
      - USERS_FILE_PATH=/app/users
    ports:
      - "1991:1991"
    volumes:
      - ./:/app/
      - ./storage:/storage
```

The `dev` Docker target uses [air](https://github.com/air-verse/air) for hot reload on `.go` and `.sql` file changes.
