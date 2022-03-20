# Parking Lot
Parking Lot is sample boilerplate usecase for workshop tokopedia academy. 

## Configuration & Dependency
- [Docker Compose](https://docs.docker.com/compose/install/) : Container manager.
- [MySQL](https://hub.docker.com/_/mysql): Used docker mysql image.
- [Adminer](https://hub.docker.com/_/adminer): Used for manage mysql on docker.

## How to run the project
### Initial config docker env configuration
```
# Copy the config env
parkinglot$ cp .env.sample .env

# Change password default mysql root on .env
MYSQL_ROOT_PASSWORD=sample

# Run docker for installation
docker-compose up
```
### Initial config docker env configuration
```
# Copy the config env
parkinglot/files/etc/parkinglot$ cp parkinglot.development.yaml.sample parkinglot.development.yaml

# Change the config parkinglot.development.yaml
version: "0.0.1"

server: 
  name: "parkinglot"
  http: 
    address: ":8080"
    write_timeout: 1
    read_timeout: 2
    max_header_bytes: 500000
    enable: true


database:
  testing: true
  dsn: "root:sample@tcp(127.0.0.1:3306)/test_parking_lot?charset=utf8mb4&parseTime=True&loc=Local"
  max_conns: 15
  max_idle_conns: 5
  max_retry: 5

new_relic:
  app_name: "Parking Lot"
  secret: "this_secret_generated" #generated from one.newrelic.com
```

### Running golang apps
```
parkinglot$ make app
```
## The endpoint list
### Get Parking List

- GET [/parking/list](http://localhost:8080/parking/list)
```
{
    "data": [
        {
            "plat_no": "DA123DEF",
            "slot_number": 1,
            "registration_date": "2022-03-12T01:49:41+08:00"
        }
    ]
}
```