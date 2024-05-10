FROM golang:1.22 as testing

RUN apt-get update && apt-get upgrade -y

RUN go install github.com/axw/gocov/gocov@latest
RUN go install github.com/AlekSi/gocov-xml@latest
RUN go install github.com/matm/gocov-html/cmd/gocov-html@latest
RUN go install mvdan.cc/gofumpt@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

RUN go install github.com/google/wire/cmd/wire@latest

WORKDIR /go/src/github.com/dqfan2012/go-scratchpad/
COPY . .
RUN mkdir -m 777 /.cache
RUN chmod -R 777 .

# FROM testing as builder
# RUN CGO_ENABLED=0 go build -ldflags '-s -w' -o go-scratchpad

# FROM scratch

# WORKDIR /
# COPY --from=builder /go/src/github.com/wpengine/dqfan2012/go-scratchpad .
# ENTRYPOINT ["/go-scratchpad"]
