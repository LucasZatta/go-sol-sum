FROM golang:alpine


ENV GIN_MODE=release
ENV PORT=8080

ADD /src 

RUN go get github.com/gin-gonic/gin




EXPOSE $PORT

ENTRYPOINT ["./app"]