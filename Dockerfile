#Base golang image
FROM golang:1.19.3-alpine

#Create directory inside the image
WORKDIR /app


#Copy the go mod and install inside the directory
COPY go.mod ./
RUN go mod download

#Copy all inside the app directory
COPY ./ ./

RUN go build -o /main ./cmd/app

EXPOSE 8080

CMD [ "/main" ]