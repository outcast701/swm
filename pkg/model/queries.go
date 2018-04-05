package model

var pQueries = map[string]string{
  "rpS1" : "SELECT id, a, w, sp, s, t, serial, uuid, muid, quid, b64 FROM %s WHERE s ORDER BY RANDOM LIMIT 1"
}
