FROM alpine:3.18.5

RUN apk add --no-cache ca-certificates

ADD karmada-webhook /bin/

CMD ["/bin/karmada-webhook"]
