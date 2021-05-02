## Overview
HTTP Backend to provide REST APIs as below
- /api/v1/accounts - To List all accounts
- /api/v1/accounts/{accounID} - To Fetch details of an account
- /api/v1/transfer - To Transfer funds from one account to another


Here we uses the Clean Code archiecture concepts to implement this service to better maintain and support this service in long run. We have defined below packages
- domain - for our entity which is account
- repository - It includes both an repository interface with limited method set and in-memory implementation for the same.
- Service - provides the methods to implement business use cases and validation
- Handler - To implement transport which in our case is HTTP

## Run the program
This program expects a command line argument to provide filename in below form if not provided it will look for accounts.json file in current working directory.
```go
go run main.go -file accounts.json 
```

## Call the APIs
Below are the curl examples to call the APIs

*Transfer Funds*
```shell
* url 'http://127.0.0.1:8080/api/v1/transfer'
* header Accept = 'application/json'
* header Content-Type = 'application/x-www-form-urlencoded'
* request {"from_account":"e68bbd7b-a3ea-4dc3-8699-0efde94a2ebd","to_account":"c2f50e66-1913-4169-bf6e-dc8c8bed6b9a","amount":1000}
* method post
```
*List Accounts*
```shell
* url 'http://127.0.0.1:8080/api/v1/accoun
```

*Get Account*
```shell
* url 'http://127.0.0.1:8080/api/v1/accounts/c2f50e66-1913-4169-bf6e-dc8c8bed6b
```

## Dependencies
This project has below two non standard library dependencies
- github.com/google/uuid v1.2.0
- github.com/gorilla/mux v1.8.0