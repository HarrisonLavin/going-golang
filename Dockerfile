FROM        golang:latest
FROM        ubuntu:latest

RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10
RUN echo "deb http://repo.mongodb.org/apt/ubuntu "$(lsb_release -sc)"/mongodb-org/3.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-3.0.list

# Update apt-get sources AND install MongoDB
RUN apt-get update && apt-get install -y mongodb-org
RUN mkdir -p /data/db
EXPOSE 27017
ENTRYPOINT ["/usr/bin/mongod"]

RUN mkdir -p /quoteGenerator

WORKDIR /quoteGenerator

ADD . /quoteGenerator

RUN go build ./quoteGenerator.go

CMD ["./quoteGenerator"]