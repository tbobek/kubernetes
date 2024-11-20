FROM golang:latest AS builder
# ARG BASE_IMAGE
# FROM ${BASE_IMAGE} as builder

WORKDIR /src
COPY . . 
RUN go get . 
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app 

# CMD ["/src/app", "-p=9876"]

FROM scratch
ENV WORKER_PORT=8765
COPY --from=builder /src/app /
CMD ["/app"]

