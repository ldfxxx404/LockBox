#!/bin/sh

set -euo pipefail

export $(grep -v '^#' .env | xargs)

DUMP_FILE="${1:-$(ls -t dump_*.dump | head -n1)}"

if [[ ! -f "${DUMP_FILE}" ]]; then
  echo "Dump file ${DUMP_FILE} not found!"
  exit 1
fi

echo "Restoring from ${DUMP_FILE}..."

docker cp "${DUMP_FILE}" "${POSTGRES_NAME}:/tmp/dump.dump"

docker exec -i "${POSTGRES_NAME}" pg_restore \
  -U "${POSTGRES_USER}" \
  -d "${POSTGRES_DB}" \
  --clean \
  --no-owner \
  --exit-on-error \
  -v /tmp/dump.dump

docker exec -t "${POSTGRES_NAME}" rm /tmp/dump.dump

echo "Restore from ${DUMP_FILE} complete"
