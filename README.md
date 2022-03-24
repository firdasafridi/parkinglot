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
  
# Change the config newrelic-infra.yml
license_key: this_secret_generated #generated from one.newrelic.com
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

### Get Nearest Empty Parking Spot
- GET [/parking/empty](http://localhost:8080/parking/empty)
```
{
    "data": 1
}
```

### Get Parking location by Plat Number
- GET [/parking?plat_no=K327714LGG](http://localhost:8080/parking?plat_no=K327714LGG)
```
{
    "data": 11
}
```

### Get Parking History by Date
- GET [/parking/history](http://localhost:8080/history?start_date=2020-01-01&end_date=2020-01-20)
```
{
  "data": [
    {
      "history_id": 1321,
      "plat_no": "O135509SOU",
      "slot_number": 983,
      "registration_date": "2020-01-01T22:56:27Z"
    },
    {
      "history_id": 1706,
      "plat_no": "N922759HQR",
      "slot_number": 941,
      "registration_date": "2020-01-01T18:54:55Z"
    }ß
  ]
}
```

### Get Report Number of Vehicles ßPer Day
- GET [/parking/history/daily-report](http://localhost:8080/parking/history/daily-report)
```
{
  "data": {
    "total_days": 367,
    "reports": [
      {
        "date": "2020-01-01",
        "total_vehicle": 113
      },
      {
        "date": "2020-01-02",
        "total_vehicle": 303
      }
  }
}
```

### Get Country Information (Upstream)
- GET [/detail/{country}](http://localhost:8080/detail/indonesia)
```
{
  "data": [
    {
      "name": {
        "common": "Indonesia",
        "official": "Republic of Indonesia",
        "nativeName": {
          "ind": {
            "official": "Republik Indonesia",
            "common": "Indonesia"
          }
        }
      },
      "tld": [
        ".id"
      ],
      "cca2": "ID",
      "ccn3": "360",
      "cca3": "IDN",
      "cioc": "INA",
      "independent": true,
      "status": "officially-assigned",
      "unMember": true,
    }
  ]
}
```
