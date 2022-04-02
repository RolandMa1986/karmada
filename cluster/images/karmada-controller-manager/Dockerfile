FROM alpine:3.13.9

RUN apk add --no-cache ca-certificates

ADD karmada-controller-manager /bin/

CMD ["/bin/karmada-controller-manager"]
