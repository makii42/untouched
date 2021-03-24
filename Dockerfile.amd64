FROM alpine
RUN apk add --no-cache git \
    && mkdir /data
WORKDIR /data
ADD dist/untouched_linux_amd64 /bin/untouched
