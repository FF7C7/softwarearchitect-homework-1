[![Build Status](https://travis-ci.org/FF7C7/softwarearchitect-homework-1.svg)](https://travis-ci.org/FF7C7/softwarearchitect-homework-1)

## Local
```bash
git clone https://github.com/FF7C7/softwarearchitect-homework-1.git
cd softwarearchitect-homework-1
docker-compose up -d
# check http://localhost:80/health/
docker-compose logs
docker-compose down
```

## DockerHub
https://hub.docker.com/r/ff7c7/softwarearchitect-homework-1
```bash
git clone https://github.com/FF7C7/softwarearchitect-homework-1.git
cd softwarearchitect-homework-1/deploy
docker-compose up -d
# check http://localhost:80/health/
docker-compose logs
docker-compose down
```
