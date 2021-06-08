#! /bin/bash
apt-get update
apt install -y docker.io #Install docker
apt install -y docker-compose
systemctl enable --now docker #Start docker and enable it to start after the system reboot
usermod -aG docker ubuntu #Give ubuntu user administrative privileges to docker

curl https://raw.githubusercontent.com/sharran-murali/go_redis_service/main/docker-compose.yml > docker-compose.yml
docker-compose up

