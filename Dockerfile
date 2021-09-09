FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src/ ./

RUN go build -o /main

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /main

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/main"]