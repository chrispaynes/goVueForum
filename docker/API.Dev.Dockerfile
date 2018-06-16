FROM golang:1.10.3
WORKDIR /go/src/goVueForum/api
COPY ./docker/runner.conf .
COPY ./api .
RUN export GOBIN="/go/bin" \
    && go get ./... \
    && go get github.com/pilu/fresh
ENTRYPOINT [ "fresh" ]