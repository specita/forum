FROM alpine:3.5
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Chongqing /etc/localtime \
    && echo "Asia/Chongqing" > /etc/timezone \
    && apk del tzdata \
    && rm -rf /var/cache/apk/*

COPY /html/index.html /html/index.html
ADD forum /forum

ENTRYPOINT [ "/forum" ]
