FROM alpine:latest as build
WORKDIR /app
COPY ./app .
RUN apk update && apk add go
RUN go get .
RUN go build -o app

FROM alpine:latest
COPY --from=build app/app .
CMD ["./app"]
