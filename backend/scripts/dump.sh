#!/bin/bash
set -euo pipefail

export $(grep -v '^#' .env | xargs)

DUMP_FILE="dump_$(date +%Y%m%d_%H%M%S).dump"

docker exec -t "${POSTGRES_NAME}" pg_dump \
  -U "${POSTGRES_USER}" \
  -d "${POSTGRES_DB}" \
  -F c \
  -f /tmp/dump.dump

docker cp "${POSTGRES_NAME}:/tmp/dump.dump" "${DUMP_FILE}"

docker exec -t "${POSTGRES_NAME}" rm /tmp/dump.dump

echo "Backup saved to ${DUMP_FILE}"
