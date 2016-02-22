# go-mathador
## Application purpose
This application has three purpose
  * provide a way to get a mathador's roll of dice, and it's answers
  * make connection between sqlite databases (20 databases) and a Redis database for speed purpose
  * it appears that if every datas were injected in the Redis db, Redis run out of ram. Must found a strategy not to run out of ram, and keep a fast response
## Dependencies
1. "github.com/gorilla/mux"
