#!/bin/bash

# Get commandline arguments
while (( "$#" )); do
  case "$1" in
    --mysql-ip)
      mysqlIp="$2"
      shift
      ;;
    --test)
      test="true"
      ;;
    *)
      shift
      ;;
  esac
done

# Common variables
mysql="mysql"Ì
password="password"
app="app"

####################
### Docker MySQL ###
####################

if [[ $test == "true" ]]; then
  # Variables
  dockerNetwork="mynetwork"
  mysqlIp="mysql"

  # Network
  docker network create \
    --driver bridge \
    $dockerNetwork

  # MySQL
  docker run \
    -d \
    --name $mysql \
    --network $dockerNetwork \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=$password \
    mysql:8

  # Wait until container successfully initializes
  sleep 4

  # App
  docker build \
    -t $app \
    ./simulator/.

  docker run \
    --rm \
    --name $app  \
    --network $dockerNetwork \
    -e MYSQL_IP=$mysqlIp \
    -e PASSWORD=$password \
    $app

  echo "Successful."
  exit 0
fi

# Check if remote MySQL IP is given
if [[ $test == "" ]] && [[ $mysqlIp == "" ]]; then
  echo "MySQL IP (--mysql-ip) should be given."
  exit 1
fi

####################
### Remote MySQL ###
####################

docker build \
  -t $app \
  ./simulator/.

docker run \
  --rm \
  --name $app  \
  -e MYSQL_IP=$mysqlIp \
  -e PASSWORD=$password \
  $app
#########

# # To test connection
# docker run --rm \
#   imega/mysql-client \
#   mysql --host=example.com --user=root --password=123321 --database=test --execute='show tables;'
