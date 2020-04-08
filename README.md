# Rest API in Go with mux Router and MySQL Database

> Simple RESTful API to create, read, update and delete employees

## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux

# Install MySQL driver
go get -u github.com/go-sql-driver/mysql
```

``` bash
go build
./Go_RestAPI
```

## Endpoints

### Get All Employees
``` bash
GET api/employees
```
### Get Single Employee
``` bash
GET api/employee/{id}
```

### Delete Employee
``` bash
DELETE api/employees/{id}
```

### Create Employee
``` bash
POST api/employees
```

### Update Employee
``` bash
PUT api/employees/{id}
