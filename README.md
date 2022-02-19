# Running in docker-compose

Complete the following one-time setup steps below, then run the program using `docker-compose up`. To rebuild the images (which is necessary for recompiling the server binary) run `docker-compose up --build`.

Prerequisites:

- install docker
- install docker-compose
- install mysql or mariadb

### Setting up MySQL/Mariadb

dbs either need to be stored at /var/lib/mysql or volume path in docker-compose.yml needs to be changed

### Environment variables

Set the following:
