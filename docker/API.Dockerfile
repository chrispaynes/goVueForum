FROM golang:1.10.3 AS build_stage
WORKDIR /go/src/goVueForm/
COPY . .
RUN export GOBIN="/go/bin" \
    && go get ./... \
    && CGO_ENABLED=0 GOOS=linux go install ./api/cmd/main.go \
    && rm -rf /go/src/goVueForm/

FROM scratch
COPY --from=build_stage /go/bin/main .
ENTRYPOINT ["./main"]