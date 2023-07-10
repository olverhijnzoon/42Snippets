FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    sysstat \
    acct \
    quota

WORKDIR /app

COPY ./resource_monitor.sh .
RUN chmod +x resource_monitor.sh