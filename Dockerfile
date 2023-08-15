FROM golang:1.21-bullseye as build

WORKDIR /go/src/app
COPY . .

RUN go mod tidy
RUN go vet . 
RUN CGO_ENABLED=0 go build -o /go/bin/app .

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /

EXPOSE 8080
CMD ["/app"]