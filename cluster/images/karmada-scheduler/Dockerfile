FROM alpine:latest

RUN apk add --no-cache ca-certificates

ADD karmada-scheduler /bin/

CMD ["/bin/karmada-scheduler"]
