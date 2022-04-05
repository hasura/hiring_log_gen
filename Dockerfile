FROM scratch

LABEL maintainer="hasura <hasura@gmail.com>"

ARG TARGETARCH

WORKDIR /app
COPY dist/hiring_log_gen_linux_$TARGETARCH/hiring_log_gen /

ENTRYPOINT ["/app/hiring_log_gen"]
