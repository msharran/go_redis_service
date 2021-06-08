# In memory key value store service 

#### **Tech Stack:** Go, Redis, Docker, Terraform, AWS
---
## Deploy server 
 >Pre-requisites: Install & Start docker; Install terraform

### Clone the repository 
```bash 
git clone https://github.com/sharran-murali/go_redis_service.git
```

### Change directory
```bash
cd go_redis_service
```

### To build, test, containerize, push and deply, run the following command
```bash
bash deploy.sh
```

> Note: Please refer [deploy.sh](./deploy.sh) to refer the steps involved

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

**NOTE:** I have pushed the image in my DockerHub Repository

https://hub.docker.com/r/sharran/go_redis_app

---
