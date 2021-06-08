# In memory key value store HTTP API Service 

#### **Tech Stack:** Go, Redis, Docker, Terraform, AWS
---

## API Docs

In memory store API
- POST `server-url`/set
- GET  `server-url`/get/:key
- GET  `server-url`/search?prefix=
- GET  `server-url`/search?suffix=

Prometheus Metrics
- GET  `server-url`/metrics

>`server-url` will be printed as terraform output

---

Deploy server 
 >Pre-requisites: Install & Start docker; Install terraform

Clone the repository 
```bash 
git clone https://github.com/sharran-murali/go_redis_service.git
```

Change directory
```bash
cd go_redis_service
```

To build, test, containerize, push and deply, run the server, use the following command
```bash
bash deploy.sh
```
---

## [deploy.sh](./deploy.sh)

```bash

#! /bin/bash

#build
echo "Building container........"
docker build -t sharran/go_redis_app . || exit

#test
echo "Running tests........"
docker-compose up -d redis || exit
docker run --net=host --rm -v "$PWD":/usr/src/app -w /usr/src/app golang:1.16 go test -run "^Test" || exit
docker-compose down

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
```

--- 

**NOTE:** I have pushed the image in my DockerHub Repository

https://hub.docker.com/r/sharran/go_redis_app

---
