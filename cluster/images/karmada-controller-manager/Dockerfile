FROM alpine:3.12.12

RUN apk add --no-cache ca-certificates

ADD karmada-controller-manager /bin/

CMD ["/bin/karmada-controller-manager"]
