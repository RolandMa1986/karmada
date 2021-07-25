FROM alpine:3.12.7

RUN apk add --no-cache ca-certificates

ADD karmada-scheduler /bin/

CMD ["/bin/karmada-scheduler"]
