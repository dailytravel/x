## Services

rm -rf account base cms community configuration finance hrm insight marketing payment sales
./init.sh account base cms community configuration finance hrm insight marketing payment sales
./g.sh account base cms community configuration finance hrm insight marketing payment sales

docker-compose up -d rabbitmq
docker-compose up -d redis