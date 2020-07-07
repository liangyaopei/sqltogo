[English Version](README.md)
## 说明
将一个创建表的SQL语句转换成Golang的ORM结构体的go函数。
例子：下面是一个创建`user`表的sql语句
```sql
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
函数`SqlToGo`将其转化成下面的代码。其中，包名是可以选择的。
```go
package data

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

## 下划线命名改为驼峰式命名
在SQL的命名规范中，字段的命名一般都是下划线分隔的,例如`ip_address`。而Golang的`struct`的字段的命名是驼峰式的。
`SqlToGo`会将其字段命名转化为驼峰式的。

## 测试运行
测试会从`input.sql`中读取SQL语句，然后打印输出，并且对输入内容进行`gofmt`，然后保存在`_output.go`
```sh
cd sql_to_go_test
go test -v .
```
