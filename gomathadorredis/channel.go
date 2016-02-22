package gomathadorredis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	redisConn redis.Conn
)

// Channel : init and return a Redis communication Channel
func Channel() {
	fmt.Print("Init Redis channel")
}
