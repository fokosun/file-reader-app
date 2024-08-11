# Use the official Alpine image as a base
FROM --platform=linux/arm64 alpine:latest

# Install OpenSSH server
RUN apk add --no-cache openssh

# Create directories for SSH host keys and user data
RUN mkdir -p /etc/ssh /var/run/sshd /home/sftpuser/upload

# Add a new user for SFTP
RUN adduser -D -h /home/sftpuser -s /sbin/nologin sftpuser

# Set a password for the SFTP user (you can change 'password' to something else)
RUN echo "sftpuser:password" | chpasswd

# Change ownership of the SFTP directory
RUN chown -R sftpuser:sftpuser /home/sftpuser

# Copy the SSH daemon configuration
RUN echo "Match User sftpuser\n\
      ChrootDirectory /home/sftpuser\n\
      ForceCommand internal-sftp\n\
      AllowTcpForwarding no" >> /etc/ssh/sshd_config

# Expose the SSH port
EXPOSE 22

# Run the SSH daemon
CMD ["/usr/sbin/sshd", "-D"]
