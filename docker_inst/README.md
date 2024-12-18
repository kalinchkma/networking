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