FROM golang

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o quiz-app-api ./cmd/main.go

CMD ["./quiz-app-api"]