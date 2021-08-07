FROM alpine:3

RUN apk add --no-cache ca-certificates

ADD karmada-webhook /bin/

CMD ["/bin/karmada-webhook"]
