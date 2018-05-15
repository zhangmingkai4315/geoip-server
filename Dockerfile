FROM golang:onbuild

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o geoip .

EXPOSE 8080
CMD [ "/app/geoip server" ]