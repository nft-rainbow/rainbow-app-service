FROM golang:1.20.8-bookworm AS build-stage
# FROM golang:1.21.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./
# set go proxy https://goproxy.cn/
ENV GOPROXY=https://goproxy.cn  
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /rainbow-apps

# copy the binary from build stage
FROM debian:bookworm AS build-release-stage
# FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /
COPY --from=build-stage /rainbow-apps /rainbow-apps

USER nonroot:nonroot
EXPOSE 8080

CMD ["/rainbow-apps"]