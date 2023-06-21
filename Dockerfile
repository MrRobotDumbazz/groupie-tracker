FROM golang:1.17-alpine3.16 AS builder
LABEL stage=builder 
ENV GO111MODULE=on
WORKDIR /app 
COPY . .
RUN go build -o main ./cmd/main.go

FROM alpine:3.16 AS runner 
LABEL stage=runner 
LABEL maintainer="Made by AmayevArtyom && Mr.RobotDumbazz"
LABEL org.label-schema.description="groupie-tracker"
WORKDIR /app
COPY --from=builder /app/main ./
COPY /static /app/static 
COPY /module /app/module
COPY /templates /app/templates
EXPOSE 8180 
CMD ["./main"]  