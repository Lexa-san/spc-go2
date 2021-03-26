FROM golang:alpine as builder 

RUN apk update && apk upgrade && apk add --no-cache git 

RUN mkdir /app
WORKDIR /app 

ENV G0111MODULE=on

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippyserver


FROM alpine:latest 

RUN apk --no-cache add ca-certificates

RUN mkdir /app 
WORKDIR /app 
COPY --from=builder /app/<your_repository_with_source_code> .

CMD ["./<your_repository_with_source_code>"]