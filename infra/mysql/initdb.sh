#!/usr/bin/env bash
set -euo pipefail

db_name="maindb"
mysql -u root -e "CREATE DATABASE IF NOT EXISTS ${db_name}"
mysql -u root ${db_name} < /schema.sql
