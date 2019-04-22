#
# Copyright (c) 2019 LG Electronics Inc.
# SPDX-License-Identifier: Apache-2.0
#

FROM golang as builder
WORKDIR /go/src/github.com/lgsvl/data-marketplace-stream-delivery
RUN go get github.com/Masterminds/glide
COPY . .
RUN glide up --strip-vendor
RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -v -a --ldflags '-w -linkmode external -extldflags "-static"' -installsuffix cgo -o bin/data-stream-delivery main.go


FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/lgsvl/data-marketplace-stream-delivery/bin/data-stream-delivery .


CMD ["./data-stream-delivery"]

EXPOSE 7778
