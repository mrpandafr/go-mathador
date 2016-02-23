package gomathadorredis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//EventListener : listen and interact with redis channel
func EventListener(psc redis.PubSubConn) {
	for {
		switch n := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("Message: %s %s\n", n.Channel, n.Data)
		case redis.PMessage:
			fmt.Printf("PMessage: %s %s %s\n", n.Pattern, n.Channel, n.Data)
		case redis.Subscription:
			fmt.Printf("Subscription: %s %s %d\n", n.Kind, n.Channel, n.Count)
			if n.Count == 0 {
				return
			}
		case error:
			fmt.Printf("error: %v\n", n)
			return
		}
	}
}
