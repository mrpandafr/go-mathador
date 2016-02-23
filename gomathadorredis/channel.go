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
	}
	c.State("open")
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

func (c *channel) ServiceKey() string {
	return c.ChannelName + "." + c.ServiceName
}
