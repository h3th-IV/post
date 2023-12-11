# Learning Postgres In Golang

- [X] zSetting Up dB

    - Login with the Default User (e.g., postgres):
        - psql -U postgres

    - create new user 
        - CREATE USER newUser WITH PASSWORD 'newPassword';
    
    - create DATABASE
        - CREATE DATABASE newDatabase;
    
    - grant permission
        - GRANT ALL PRIVILEGES ON DATABASE mydatabase TO newUser;
        - [optional] logout, login with new user to new databse
        - psql -U your_username -d your_database
    
    - environment variables will look like this:
        - PG_HOST=localhost
        - PG_PORT=5432
        - PG_USER=newUser
        - PG_PASSWORD=newPassword
        - PG_DBNAME=newDatabase

