# Run using docker compose

Complete the following one-time setup steps below, then run the program using `docker-compose up`. To rebuild the images (which is necessary for recompiling the server binary) run `docker-compose up --build`.

# Environment variables

Create a .env file in the root of the app with the following keys:
```
MARIADB_ROOT_PASSWORD=${whatever you want}
DATABASE_NAME=${whatever you want}
DATABASE_USERNAME=${whatever you want}
DATABASE_PASSWORD=${whatever you want}
```
The values can be whatever you want and will be used when the MariaDB image is built to set the credentials.

# Volumes

A database volume will be stored at /var/lib/mysql/docker. The database will be instantiated according to the script ```/db/docker-entrypoint-initdb.d/schema.sql``` only upon first build. Any subsequent changes to the credentials or the schema will require any files created by Docker at the mount destination to be deleted before the changes can take effect.
