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
