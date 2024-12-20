# Docker 

Docker container virtualize at the operating system level, where `Virtial machines` virtualize hardware, they emulate what a physical computer donse at a very low level

###### Images vs Containers
- A "docker image" is the read-only definition of a container
- A "docker container" is a virtialized read-write environtment

See list of images
```bash
docker images
```

### Basic commands

Run docker with port export
```bash
docker run -d -p <main_machine_port>:<docker_container_port> <docker_image>
```

Stopping a container
```bash
docker stop CONTAINER_ID
```

Stopping container by issuing a `SIGKILL` signal to the container
```bash
docker kill CONTAINER_ID
```

Exec
```bash
docker exec CONTAINER_ID <command>
```

Find the process
```bash
docker exec CONTAINER_ID netstat -ltnp
```

Live Shell
```bash
docker exec -it CONTAINER_ID /bin/sh
```
- -i makes the `exec` command interactive
- -t gives us a tyy(keyboard) interface
- `/bin/sh` or `sh` is the path command we are running

See the dokcer running process
```bash
docker stats
```

Volumes

Create volumes
```bash
docker volume create volume_name
```

List all available volume
```bash
docker volume ls
```

Inspect volume
```bash
docker volume inspect volume_name
```

Pull docker image from docker hub
```bash
docker pull docker_image_name
```

Ex:
```bash
docker pull ghost
```

Start container
```bash
docker run -d -e NODE_ENV=development -e url=http://localhost:3001 -p 3001:2368 -v demo_volume:/var/lib/ghost/content ghost
```
- `-e NODE_ENV=development` set the environtment variable
- `-e url=http://localhost:3001` set another environtment variable
- `-p 3001:2368` port forwarding
- `-v demo_volume:/var/lib/ghost/content` mounts the `demo_volume` to the `/var/lib/ghost/content` path in the container

Restart running container
```bash
docker restart CONTAINER_ID
```

Deleting the volume
```bash
docker volume rm VOLUME_NAME
```

Rin containger without network access
```bash
docker run -d --network none IMAGE
```

Create network
```bash
docker network create NETWORK_NAME
```

List available network
```bash
docker network ls
```

Ex: load balancing using `caddy`

```bash
docker pull caddy
```
Create demo file to separete server context

run server 1
```bash
docker run -d -p 8000:80 -v $PWD/index1.html:/usr/share/caddy/index.html caddy
```
run server 2
```bash
docker run -d -p 8001:80 -v $PWD/index2.html:/usr/share/caddy/index.html caddy
```
run server with name specified in same network
create networl
```bash
docker network create ganja
```

```bash
docker run -d --network ganja --name caddy1 -p 8000:80 -v $PWD/index1.html:/usr/share/caddy/index.html caddy
```
```bash
docker run -d --network ganja --name caddy2 -p 8001:80 -v $PWD/index2.html:/usr/share/caddy/index.html caddy
```

```bash
docker run -it --network ganja docker/getting-started /bin/sh
```

create `Caddyfile`

run loadbalancer
```bash
docker run -d --network ganja -p 8080:80 -v $PWD/Caddyfile:/etc/caddy/Caddyfile caddy
```

#### Dockerfile

```code
# Use base image
FROM debian:stable-slim

# execute the command
CMD ["echo", "hello world"]
```
Save as a `Dockerfile` and build the image
```bash
docker build . -t image_name:tag
```
- `-t image_name:tag` to specify image name and tag

```bash
docker build . -t test:latest
```
```bash 
docker run test
```




