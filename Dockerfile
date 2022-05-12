FROM golang:alpine AS build
RUN apk add git
WORKDIR /build

# mod files go first to prevent unchanged downloads
COPY go.* ./
RUN go mod download
RUN go mod verify

# copy go files and build
COPY . .
RUN go build -o grencon github.com/thecodingcouple/gren-con-api/cmd/grencon


FROM alpine
RUN adduser -S -D -H -h /api apiuser
USER apiuser
COPY --from=build /build/grencon /api/
WORKDIR /api
CMD [ "./grencon" ]
