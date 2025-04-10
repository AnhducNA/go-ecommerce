FROM golang:alpine AS builder

WORKDIR /builder

COPY . .

RUN go mod download

RUN go build -o crm.shopdev.com ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /builder/crm.shopdev.com /

ENTRYPOINT [ "/crm.shopdev.com", "config/local.yaml" ]