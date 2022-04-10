FROM alpine:3.14.6

RUN apk add --no-cache ca-certificates

ADD karmada-controller-manager /bin/

CMD ["/bin/karmada-controller-manager"]
