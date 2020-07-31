[![Go Report Card](https://goreportcard.com/badge/github.com/liangyaopei/sqltogo)](https://goreportcard.com/report/github.com/liangyaopei/sqltogo)
[![GoDoc](https://godoc.org/github.com/liangyaopei/sqltogo?status.svg)](http://godoc.org/github.com/liangyaopei/sqltogo)
[中文版说明](./README_zh.md)

## Description
This repository provide a way to convert SQL create statement to Golang struct (ORM struct) by parsing the SQL statement.
For example, with the input
```mysql
CREATE TABLE `USER`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    `ip_address` INT  NOT NULL DEFAULT 0 COMMENT 'ip_address',
    `nickname`    VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'user note',
    `description` VARCHAR(256) NOT NULL DEFAULT '' COMMENT 'user description',
    `creator_email` VARCHAR(64) NOT NULL DEFAULT '' COMMENT 'creator email',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
    `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete time',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='user table';
```
the function `SqlToGo` will convert it to as follow, with the `package` name is optional.
```go
package sql_to_go_test

import (
	"time"
)

type USER struct {
	Id           uint      `comment:"primary key"`
	IpAddress    int       `comment:"ip_address"`
	Nickname     string    `comment:"user note"`
	Description  string    `comment:"user description"`
	CreatorEmail string    `comment:"creator email"`
	CreatedAt    time.Time `comment:"create time"`
	DeletedAt    time.Time `comment:"delete time"`
}
```

## Snake case To Camel case
In SQL, the naming convention of a filed is snake case, such as `ip_address`, while in Golang, 
the naming convention of struct's field is Camel case. So the `SqlToGo` function will convert snake case to
Camel case.

## Test & run
The test will read the sample sql from `input.sql`, and print the output, then `gofmt` it and save
the result in `_output.go`
```sh
cd sql_to_go_test
go test -v .
```