package model

import (
  "fmt"

  "github.com/outcast701/swm/pkg/config"
)

var pQueries = map[string]string{
  "rpS1" : "SELECT id, a, w, sp, s, t, serial, uuid, muid, quid, b64 FROM %s WHERE s ORDER BY RANDOM LIMIT 1",
  "cP" : "CREATE TABLE %s (id BIGSERIAL, a INTEGER, w INTEGER, sp INTEGER, s BOOLEAN, t VARCHAR(1), serial BIGINT, uuid VARCHAR(38), muid VARCHAR(38), quid VARCHAR(38), b64 VARCHAR(16))",
}
var conf config.Config

func init() {
    conf, _ = config.LoadConfig("~/.swm/config.json")

}

func GetQuery(query string) string {
  return fmt.Sprintf(pQueries[query], conf.TableName)
}
