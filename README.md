# groupie-tracker
# build:
    docker build --rm -t groupie-tracker .
    docker images

# run:
	docker container run -p 8080:8080 --detach --name groupie-tracker groupie-tracker
    docker ps -a

# check file system:
    docker exec -it ascii-art-web /bin/bash ### may not work use sh instead of bash 
    ls -l
    exit

# check docker metadata:
    docker inspect \
        --format '{{ index .Config.Labels "maintainer"}}' \
    ascii-art-web
