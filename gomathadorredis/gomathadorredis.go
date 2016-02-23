package gomathadorredis

// Error represents an error returned in a command reply.
type Error string

func (err Error) Error() string { return string(err) }

// Channel represents a cummication channnel
type Channel interface {
	State(state string) (reply bool, err error)
	ServiceKey() string
}
