FROM golang AS builder

COPY lbhelper.go .
RUN CGO_ENABLED=0 go build lbhelper.go

FROM scratch
COPY --from=builder /go/lbhelper /
CMD ["/lbhelper"]