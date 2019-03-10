#!/usr/bin/env sh

set -eu

# Initialize
BASE_DIR=$(dirname $0)
mysql -h ${DB_HOST} -u${DB_USER} -p${DB_PW} ${DB_NAME} < ${BASE_DIR}/query.sql
