package gomathadorredis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/soveran/redisurl"
)

type channel struct {
	ServiceName string
	ChannelName string
	state       string
	conn        redis.Conn
	Psc         redis.PubSubConn
}

// NewChannel : init and return a communication Channel
func NewChannel(channelname string, servicename string) (Channel, error) {
	redisConn, err := redisurl.Connect()
	if err != nil {
		return nil, err
	}
	c := &channel{
		ServiceName: servicename,
		ChannelName: channelname,
		state:       "online",
		conn:        redisConn,
		Psc:         redis.PubSubConn{Conn: redisConn},
	}
	c.Psc.Subscribe(c.ServiceKey())
	c.State("open")
	c.Publish("open")
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

func (c *channel) Publish(data string) error {
	redisConn, err := redisurl.Connect()
	if err != nil {
		return err
	}
	redisConn.Send(c.ServiceKey(), data)
	redisConn.Flush()
	for {
		_, err := redisConn.Receive()
		if err != nil {
			return err
		}
		// process pushed message
	}
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
