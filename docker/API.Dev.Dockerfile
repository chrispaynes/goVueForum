FROM golang:1.11.0
WORKDIR /go/src/goVueForum/api
COPY ./api .
RUN export GOBIN="/go/bin" \
    && go get ./... \
    && go get github.com/pilu/fresh
ENTRYPOINT [ "fresh" ]