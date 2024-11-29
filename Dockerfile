# Dockerfile
FROM ubuntu:latest

# Install SSH and net-tools (for ping)
RUN apt-get update && apt-get install -y \
    openssh-server \
    net-tools \
    iputils-ping && \
    mkdir /var/run/sshd

# Set root password
RUN echo "root:password" | chpasswd

# Allow SSH login with root
COPY ssh_config/sshd_config /etc/ssh/sshd_config

# Expose SSH port
EXPOSE 22

# Create a data directory
RUN mkdir -p /data && chmod 777 /data

# Start SSH service
CMD ["/usr/sbin/sshd", "-D"]
