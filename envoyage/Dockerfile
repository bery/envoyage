# Dockerfile

FROM golang:latest as base

WORKDIR /app

COPY go.* ./

RUN go mod download
RUN go mod verify

COPY *.go ./

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /go/app/binary .

CMD ["air"]


FROM golang:bookworm

WORKDIR /app

COPY --from=base /go/app/binary /app

EXPOSE 8000

CMD [ "/app/binary" ]