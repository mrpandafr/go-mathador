package gomathadorredis

import (
	"fmt"
	"sync"

	"github.com/garyburd/redigo/redis"
	"github.com/soveran/redisurl"
)

type channel struct {
	ServiceName string
	ChannelName string
	state       string
	conn        redis.Conn
}

// NewChannel : init and return a communication Channel
func NewChannel(channelname string, servicename string) (Channel, error) {
	redisConn, err := redisurl.Connect()
	if err != nil {
		return nil, err
	}
	defer redisConn.Close()
	var wg sync.WaitGroup
	wg.Add(2)
	psc := redis.PubSubConn{Conn: redisConn}

	c := &channel{
		ServiceName: servicename,
		ChannelName: channelname,
		state:       "online",
		conn:        redisConn,
	}
	c.State("open")

	psc.Subscribe(c.ServiceKey())
	go func() {
		defer wg.Done()
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
	}()

	// This goroutine manages subscriptions for the connection.
	go func() {
		defer wg.Done()

		psc.Subscribe("example")
		psc.PSubscribe("p*")

		// The following function calls publish a message using another
		// connection to the Redis server.
		publish("example", "hello")
		publish("example", "world")
		publish("pexample", "foo")
		publish("pexample", "bar")

		// Unsubscribe from all connections. This will cause the receiving
		// goroutine to exit.
		psc.Unsubscribe()
		psc.PUnsubscribe()
	}()

	wg.Wait()
	return c, nil
}

//State : set a service state
func (c *channel) State(state string) (bool, error) {
	val, err := c.conn.Do("SET", c.ServiceKey(), state, "NX", "EX", "120")
	if err != nil {
		return false, err
	}
	if val == nil {
		fmt.Println(c.ServiceName + " already " + state)
		return false, nil
	}
	return true, nil
}

func publish(channel, data string) error {
	c, err := redisurl.Connect()
	if err != nil {
		return err
	}
	c.Send(channel, data)
	c.Flush()
	for {
		_, err := c.Receive()
		if err != nil {
			return err
		}
		// process pushed message
	}
}

func (c *channel) ServiceKey() string {
	return c.ChannelName + "." + c.ServiceName
}
