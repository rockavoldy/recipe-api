FROM golang:1.21-bullseye as build

WORKDIR /go/src/app
COPY . .

ENV RECIPE_DB_HOST "$RECIPE_DB_HOST"
ENV RECIPE_DB_USER "$RECIPE_DB_USER"
ENV RECIPE_DB_PASSWORD "$RECIPE_DB_PASSWORD"
ENV RECIPE_DB_NAME "$RECIPE_DB_NAME"
ENV RECIPE_DB_PORT "$RECIPE_DB_PORT"
ENV RECIPE_HTTP_PORT "$RECIPE_HTTP_PORT"
ENV CGO_ENABLED 0

RUN go mod download
RUN go vet . 
RUN go build -o /go/bin/app .

FROM gcr.io/distroless/static-debian11

ENV RECIPE_DB_HOST "$RECIPE_DB_HOST"
ENV RECIPE_DB_USER "$RECIPE_DB_USER"
ENV RECIPE_DB_PASSWORD "$RECIPE_DB_PASSWORD"
ENV RECIPE_DB_NAME "$RECIPE_DB_NAME"
ENV RECIPE_DB_PORT "$RECIPE_DB_PORT"
ENV RECIPE_HTTP_PORT "$RECIPE_HTTP_PORT"

COPY --from=build /go/bin/app /

EXPOSE 8080
CMD ["/app"]