#!/bin/bash

apk add --no-cache --virtual mysql-client

echo "Waiting for mysql"
until mysql -hdb -P3306 -ustatus_app -pstatus_password &> /dev/null
do
  printf "."
  sleep 1
done

echo -e "\nmysql ready"

/app/status