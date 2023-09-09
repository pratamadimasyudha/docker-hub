FROM alpine:3.18.3

RUN addgroup -g 10001 \
             -S nonroot && \
    adduser  -u 10000 \
             -G nonroot \
             -h /home/nonroot \
             -S nonroot

RUN echo -e "https://dl-cdn.alpinelinux.org/alpine/v$(cut -d'.' -f1,2 /etc/alpine-release)/main/\n \
             https://dl-cdn.alpinelinux.org/alpine/v$(cut -d'.' -f1,2 /etc/alpine-release)/community/\n \
             https://dl-cdn.alpinelinux.org/alpine/edge/testing/\n \
             https://dl-cdn.alpinelinux.org/alpine/edge/community/\n \
             https://dl-cdn.alpinelinux.org/alpine/edge/main/" > /etc/apk/repositories && \
    sed -i 's/^[ \t]*//;s/[ \t]*$//' /etc/apk/repositories

RUN apk update && \
    apk add --no-cache tini=0.19.0-r1 \
                       bind-tools=9.18.16-r0 \
                       go=1.20.7-r0