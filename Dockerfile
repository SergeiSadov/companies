FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .

RUN go mod download
COPY . .

RUN go build -o /companies cmd/companies/*.go

EXPOSE 3000

CMD [ "/companies" ]