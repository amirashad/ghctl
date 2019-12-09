FROM alpine
LABEL maintainer="amirjanov@gmail.com" 

RUN apk add --no-cache curl

COPY ./ghctl /bin/
COPY ./yq/yq /bin/

ENTRYPOINT ["/bin/ghctl"]
