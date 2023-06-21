# groupie-tracker
# build:
    docker build --rm -t ascii-art-web .
    docker images

# run:
	docker container run -p 8080:8080 --detach --name ascii-art-web ascii-art-web
    docker ps -a

# check file system:
    docker exec -it ascii-art-web /bin/bash ### may not work use sh instead of bash 
    ls -l
    exit

# check docker metadata:
    docker inspect \
        --format '{{ index .Config.Labels "maintainer"}}' \
    ascii-art-web
