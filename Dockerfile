FROM golang:1.22-alpine3.19 AS buildStage
WORKDIR /auth-svc
COPY . ./
RUN go mod download
RUN go build -o ./auth-svc ./cmd

FROM scratch AS release-stage 
WORKDIR /
COPY --from=buildStage /auth-svc/auth-svc /auth-svc
COPY --from=buildStage /auth-svc/dev.env /

EXPOSE 50051

ENTRYPOINT [ "/auth-svc" ]