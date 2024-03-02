# docker/localstack/Dockerfile
FROM localstack/localstack

COPY docker/localstack/init.sh /etc/localstack/init/ready.d/init.sh
RUN chmod +x /etc/localstack/init/ready.d/init.sh