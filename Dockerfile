FROM golang

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN make build
EXPOSE 3000
CMD ["make", "run"]
