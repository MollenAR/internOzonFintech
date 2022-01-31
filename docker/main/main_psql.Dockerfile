FROM golang:alpine as build

COPY ./ /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o main.out ./cmd/main.go

FROM alpine

COPY --from=build /app/main.out /
COPY --from=build /app/configs/. /configs/
RUN chmod +x main.out

ENTRYPOINT ["./main.out"]