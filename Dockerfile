FROM golang

RUN mkdir /app
ADD . /app/
WORKDIR /app/cmd/projectmanager
RUN go build -o projectmanager .

CMD ["./projectmanager", "-c" ,"/app/configs/local.yaml"]