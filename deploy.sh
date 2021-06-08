#! /bin/bash

#build
echo "Running build........"
go get -d -v ./...  || exit
go build -v -o ~/go/bin/goredisbin  || exit

#test
echo "Running tests........"
docker-compose up -d redis || exit
go test -v -run "^Test" || exit

#build and push image
echo "Building container........"
docker build -t sharran/go_redis_app . || exit
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