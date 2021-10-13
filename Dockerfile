FROM golang:1.16-alpine3.13

WORKDIR app/user

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /main-build

ENV DB_HOST=host.docker.internal
ENV REDIS_URL=host.docker.internal:6379

EXPOSE 8080

CMD [ "/main-build" ]