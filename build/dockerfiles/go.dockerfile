FROM golang:1.21-bullseye as BUILDER

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /runner ./cmd/services/cars-runner/main.go

FROM golang:1.21-bullseye

COPY --from=BUILDER /runner /runner

RUN mkdir /project
RUN cd /project && go mod init project

RUN apt-get  update
RUN apt-get install coreutils
