# syntax=docker/dockerfile:1

FROM golang:1.19 as BUILDER

#Active le comportent de module indépendant
ENV GO111MODULE=on

ENV CGO_ENABLED=0
ENV GOOS=$GOOS
ENV GOARCH=$GOARCH

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY ./app .
RUN go mod download \ 
&& go mod verify \
&& go build -o /build/buildedApp main/main.go

FROM scratch as FINAL

WORKDIR /main
COPY --from=BUILDER /build/buildedApp .

ENTRYPOINT ["./buildedApp"]

