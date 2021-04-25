# builder image
FROM golang:1.16-alpine as BUILDER

RUN mkdir /build

ADD /application /build/

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /firstgo_app .


# generate clean, final image for end users
FROM alpine:3.11.3

COPY --from=BUILDER ./firstgo_app .

# executable
ENTRYPOINT [ "./firstgo_app" ]
