FROM alpine:3.17

RUN apk add --no-cache ca-certificates

ADD karmada-controller-manager /bin/

CMD ["/bin/karmada-controller-manager"]
