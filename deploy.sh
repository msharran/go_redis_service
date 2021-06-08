#! /bin/bash

#build
echo "Building container........"
docker build -t sharran/go_redis_app . || exit

#test
echo "Running tests........"
docker-compose up -d redis || exit
docker run --net=host --rm -v "$PWD":/usr/src/app -w /usr/src/app golang:1.16 go test -run "^Test" || exit

#push
echo "Pushing container........"
docker push sharran/go_redis_app  || exit

#deploy
echo "Provisioning server........"
cd terraform || exit
terraform init  || exit
terraform validate  || exit
terraform plan || exit
terraform apply -auto-approve
echo "Successfully deployed........"
echo "It will take around 2 to 5 mins for the server to be up and running. Please wait..."