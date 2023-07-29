# TransferSystem

- Create .env file
- Copy the content of .env.example file to .env
- Run 'docker-compose up -d' to connect to my sql database
- Run 'docker build -t  go:1.0 . --network="host"' to build docker image
- then use http://127.0.0.1:8080/api/login with {"name":"Browsecat", "password":123456} in (raw body ) to login and this api will generate "token" that will be used in the next apis for the autherization  as Barear token
- then use http://127.0.0.1:8080/api/user 'with header authorization:token' to list all users in the system (token will be the generated token from login api)
- use http://127.0.0.1:8080/api/credit/tranferr {"transferer":"ID from users ids", "transferee":"ID from users ids", "amount":200}
