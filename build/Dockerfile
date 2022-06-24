FROM golang:1.17.3 as builder

WORKDIR /go/src/fake.com/padel-api

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/padel-api ./cmd/api

FROM scratch as scratch
COPY --from=builder /bin/padel-api /bin/padel-api
COPY --from=builder /go/src/fake.com/padel-api/config ./config
COPY --from=builder /go/src/fake.com/padel-api/internal/templates/resources /internal/templates/resources/
COPY --from=builder /go/src/fake.com/padel-api/docs /docs

ENTRYPOINT ["/bin/padel-api"]