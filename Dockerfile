FROM docker.io/ubuntu:22.04 AS build

RUN apt-get update --yes
RUN apt-get install golang ca-certificates --yes

ADD ./ /service_go
WORKDIR /service_go
ENV CGO_ENABLED=0
RUN go build -o /kate /service_go

FROM gcr.io/distroless/static

COPY --from=build /kate /kate

ENTRYPOINT /kate
