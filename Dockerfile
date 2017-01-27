FROM golang
ADD . /go/src/github.com/claytoncheung/balancebot
RUN go get github.com/bwmarrin/discordgo \
    && go install github.com/claytoncheung/balancebot
ENTRYPOINT /go/bin/balancebot
