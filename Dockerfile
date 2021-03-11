FROM  golang:1.16.0-alpine3.13 AS build

WORKDIR /berlin
COPY . .
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"'

RUN go build -o /out/berlin .


FROM  golang:1.16rc1-alpine AS bin
WORKDIR /app
COPY --from=build /out/berlin /app/berlin
ENTRYPOINT ["/app/berlin"]
