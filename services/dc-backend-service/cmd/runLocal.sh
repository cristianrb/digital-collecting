#!/bin/sh

export DCUSERSAPIPORT="8080"
export DCBACKENDAPIPORT="8081"

export PGUSER="dc_admin"
export PGPASSWORD="dc_password"
export PGHOST="localhost"
export PGPORT="5432"
export PGDATABASE="digital_collecting"

export JWTDOMAIN="com.github.cristianrb.dc-users-service"
export JWTAUDIENCE="dc-users-service"
export JWTSECRET="jwt-secret"
export JWTISSUER="digital-collecting"
export JWTREALM="test"

go run main.go