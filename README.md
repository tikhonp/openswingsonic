# openswingsonic

Open Subsonic API translation layer for Swing Music.

# dev

```yaml
services:
  oswingsonic:
    container_name: oswingsonic
    build: 
      context: .
      target: prod
      args:
        APP_VERSION: dev
    environment:
      - DEBUG=true
      - SWINGMUSIC_BASE_URL=swingmusic:1970
      - PUBLIC_SWINGMUSIC_URL=https://m...
      - CRED_PROVIDER=file
      - USERS_FILE_PATH=users
    dns:
      - 8.8.8.8
      - 1.1.1.1
    ports:
      - "1991:1991"
    volumes:
      - "./:/app/"
      - "./storage:/storage"

```
