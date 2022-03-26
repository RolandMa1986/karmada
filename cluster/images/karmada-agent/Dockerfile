FROM alpine:3.13.8

RUN apk add --no-cache ca-certificates

ADD karmada-agent /bin/

CMD ["/bin/karmada-agent"]
