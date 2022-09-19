FROM golang:latest
# 
WORKDIR /app
COPY . .
# 
RUN go build -o main main.go

EXPOSE 8090

# 
CMD [ "/app/main" ]