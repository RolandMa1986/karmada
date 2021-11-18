FROM alpine:3.13.7

RUN apk add --no-cache ca-certificates

ADD karmada-agent /bin/

CMD ["/bin/karmada-agent"]
