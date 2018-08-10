FROM alpine
RUN apk add --no-cache git \
    && mkdir /data
WORKDIR /data
ADD untouched /bin
