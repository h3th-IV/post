# Learning Postgres In Golang

- [X] zSetting Up dB

    - login: sudo -u postgres psql  or psql -U postgres

    - create DATABASE
        - CREATE DATABASE newDatabase;
    
    - create new user for  new DATABASE
        - CREATE USER newUser WITH PASSWORD 'newPassword';
    
    - grant permission
        - GRANT ALL PRIVILEGES ON DATABASE mydatabase TO newUser;
    
    
    - environment variables will look like this:
        - PG_HOST=localhost
        - PG_PORT=5432
        - PG_USER=newUser
        - PG_PASSWORD=newPassword
        - PG_DBNAME=newDatabase
