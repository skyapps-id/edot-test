# Edot Test

### Architecture
![Alt text](arc.jpeg "profile-service")

### Tech Stack
- Echo v4
- Postgresql
- Jwt
- Machinery for worker

### Quick Installation Database
1. Run all database service
    ```sh
    $ docker compose up -d
    ```
2. Restore all file migration.sql in folder per-service databse

### Run Service
1. Run service user-service
    ```sh
    $ go run main.go
    ```
2. Run service product-service
    ```sh
    $ go run main.go
    ```
3. Run service shop-warehouse-service
    ```sh
    $ go run main.go
    ```
4. Run service order-service

    terminal 1
    ```sh
    $ go run main.go server
    ```
    terminal 2
    ```sh
    $ go run main.go worker
    ```

### Test API On Insomnia 
1. Import file Insomnia_2025-05-18.json


### Contact
https://www.linkedin.com/in/aji-indra-jaya

License
----

MIT