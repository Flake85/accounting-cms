package flags

import "flag"

var host = "localhost"
var port = "54321"
var user = "cms"
var password = "cms"
var dbName = "cms"

var Host = flag.String("POSTGRES_HOST", host, "specifies postgres host")
var Port = flag.String("POSTGRES_PORT", port, "specifies postgres port")
var User = flag.String("POSTGRES_USER", user, "specifies postgres user")
var Password = flag.String("POSTGRES_PASSWORD", password, "specifies postgres password")
var DbName = flag.String("POSTGRES_DB", dbName, "specifies postgres db name")
