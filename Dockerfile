FROM golang as builder

ADD . /opt/app
WORKDIR /opt/app/
RUN CGO_ENABLED=0 go build -o /desafio-ascan .

FROM alpine
RUN apk --no-cache add curl
COPY --from=builder /desafio-ascan /desafio-ascan
WORKDIR /
ENTRYPOINT ./desafio-ascan