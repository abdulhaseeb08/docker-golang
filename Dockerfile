FROM golang:latest 

RUN mkdir /build 
WORKDIR /build

RUN export GO111MODULE=on 
RUN go get github.com/abdulhaseeb08/docker-golang/main
RUN cd /build && git clone https://github.com/abdulhaseeb08/docker-golang.git

RUN cd /build/docker-golang/main && go build 

EXPOSE 8080

ENTRYPOINT [ "/build/docker-golang/main/main"]