#!/bin/bash

set -e 

docker cp dump.dump lockbox-postgres:/tmp/dump.dump

docker exec -i lockbox-postgres pg_restore \
  -U postgres \
  -d lock_box \
  --clean \
  --no-owner \
  --exit-on-error \
  -v /tmp/dump.dump
