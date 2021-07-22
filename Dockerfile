FROM golang:alpine


ENV GIN_MODE=release
ENV PORT=8080

ADD /src /go/src/github.com/go-sol-sum

RUN go get github.com/gin-gonic/gin




EXPOSE $PORT

ENTRYPOINT ["./app"]