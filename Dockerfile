FROM ubuntu:latest
LABEL authors="kevin.ninh"

ENTRYPOINT ["top", "-b"]