#!/bin/bash

docker exec -t lockbox-postgres pg_dump \ 
  -U postgres \ 
  -d lock_box \
  -F c \
  -f /tmp/dump.dump

docker cp lockbox-postgres:/tmp/dump.dump dump.dump

docker exec -t lockbox-postgres rm /tmp/dump.dump
