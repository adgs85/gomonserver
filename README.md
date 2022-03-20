# gomonserver
docker compose:

```
./build.sh
./dockerbuild.sh
docker-compose up
```







Postgresql server required
docker-compose:
```
version: '3.1'

services:

  db:
    image: postgres
    restart: always
    environment:
      POSTGRES_PASSWORD: example
    ports:
      - "5432:5432"

```
      


DDL:
```
create database stats_db;

create table STATS_SNAPSHOT(
	host_name varchar(255),
	stat_type varchar(255),
	instance_name varchar(255),
	collected_ts timestamp,
	last_updated timestamp,
	polling_rate_ms int,
	payload text
);

create table STAT_HISTORY(
	host_name varchar(255),
	stat_type varchar(255),
	instance_name varchar(255),
	collected_ts timestamp,
	last_updated timestamp,
	polling_rate_ms int,
	payload text
);

ALTER TABLE stats_snapshot ADD PRIMARY KEY (host_name , stat_type);
```
