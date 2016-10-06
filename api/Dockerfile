FROM golang:alpine
ADD . /go/src/code.olipicus.com/go_line_bot_api
WORKDIR /go/src/code.olipicus.com/go_line_bot_api
RUN go get github.com/line/line-bot-sdk-go/linebot
RUN go build -o api .
CMD ["./api"]
