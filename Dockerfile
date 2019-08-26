FROM alpine

RUN apk add --no-cache ca-certificates && apk update
# Change TimeZone
RUN apk add --update tzdata
ENV TZ=Asia/Bangkok
# Clean APK cache
RUN rm -rf /var/cache/apk/*

ADD happy-q-api /happy-q-api
ADD .env /.env
CMD ["./happy-q-api"]
