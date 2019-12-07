FROM alpine

RUN apk add --no-cache curl

COPY ./ghctl ./

ENTRYPOINT ["./ghctl"]
