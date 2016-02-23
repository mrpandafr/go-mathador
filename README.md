# go-mathador
## Application purpose
This application has three purpose
  * provide a way to get a mathador's roll of dice, and it's answers
  * make connection between sqlite databases (20 databases) and a Redis database for speed purpose
  * it appears that if every datas were injected in the Redis db, Redis run out of ram. Must found a strategy not to run out of ram, and keep a fast response

## Roadmap
  * initial dummy REST server OK
  * redis channel OK
  * multi redis channel OK

## TODO
  * feed the DB
  * retrieve the datas

## Dependencies
### Externals
#### Currents
1. "github.com/gorilla/mux"
2. "github.com/garyburd/redigo/redis"
3. "github.com/soveran/redisurl"

#### Futures
1. "github.com/mattn/go-sqlite3"

### Cores
#### Currents
1. "fmt"
2. "html"
3. "log"
4. "net/http"
5. "encoding/json"
6. "time"

#### Futures
1. "database/sql"
2. "flag"
3. "os"
