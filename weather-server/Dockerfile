FROM golang:1.21 AS server-build
LABEL stage=build

WORKDIR /weather-server/src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 go build -v -o /weather-server/bin/ ./...


FROM redhat/ubi9-micro AS server-dist
LABEL stage=dist

WORKDIR /

COPY --from=server-build /weather-server/bin/weather-server /

ENV DEBUG=false
ENV HTTP_PORT=8080

EXPOSE 8080
ENTRYPOINT ["/weather-server"]
