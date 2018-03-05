FROM scratch
WORKDIR /docker-go
COPY docker-go .
COPY docker-go.yaml .
EXPOSE 8080
ENTRYPOINT ["/docker-go/docker-go"]