package configs

import (
	"os"
	"strconv"
)

var PG_HOST = os.Getenv("PG_HOST")
var PG_PORT = os.Getenv("PG_PORT")
var POSTGRES_DB = os.Getenv("POSTGRES_DB")
var POSTGRES_USER = os.Getenv("POSTGRES_USER")
var POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
var HTTP_PORT = os.Getenv("HTTP_PORT")
var MIGRATE_DOWN, _ = strconv.ParseBool(os.Getenv("MIGRATE_DOWN"))
var MIGRATIONS_PATH = os.Getenv("MIGRATIONS_PATH")
var TEMPLATES_PATH = os.Getenv("TEMPLATES_PATH")
var MIGRATE_VERSION_FORCE, _ = strconv.Atoi(os.Getenv("MIGRATE_VERSION_FORCE"))
var KAFKA_BROKER = os.Getenv("KAFKA_BROKER")
var SESSION_SECRET = os.Getenv("SESSION_SECRET")

var CERT_PATH = os.Getenv("CERT_PATH")
var PROBLEMS_GRPC_PORT = os.Getenv("PROBLEMS_GRPC_PORT")
var PROBLEMS_SERVICE = os.Getenv("PROBLEMS_SERVICE")
var PROBLEMS_CERT_NAME = os.Getenv("PROBLEMS_CERT_NAME")
var PROBLEMS_CERTIFICATE = CERT_PATH + PROBLEMS_CERT_NAME