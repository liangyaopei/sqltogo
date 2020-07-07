## 说明
将一个创建表格的SQL语句转换成Golang model的go语句。
根据SQL语句的model转换成对应Golang的ORM的结构体。

例子：
下面是一个创建`user`表的sql语句
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
转换成对应的orm的go文件
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

## 测试运行
`sql_to_go_test.go`文件中，读取`./data`下的`input.sql`文件，将其输出到`./data/putput.go`.
测试命令
```sh
go test -v .
```
