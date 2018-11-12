FROM debian
RUN apt-get update
RUN apt-get install -y ca-certificates
COPY go-school /
WORKDIR /
CMD ["/go-school"]
