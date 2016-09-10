FROM golang:1.6.0

ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

WORKDIR /go/src/github.com/Yicwif/CMS
ADD . /go/src/github.com/Yicwif/CMS

EXPOSE 8999

ENV SERVICE_NAME=CMS

RUN go build

ENTRYPOINT ["/go/src/github.com/Yicwif/CMS/CMS"]

