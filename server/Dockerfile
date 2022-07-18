FROM golang
MAINTAINER Stepan K. <xamust@gmail.com>
WORKDIR /server/
VOLUME ["/opt/server"]
COPY . ./
RUN make build
EXPOSE 8080
CMD [ "build/server" ]
