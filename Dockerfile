FROM debian:latest
RUN mkdir /app
ADD config/ /app
ADD main /app
WORKDIR /app
CMD ["./main"]