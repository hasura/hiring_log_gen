FROM scratch

LABEL maintainer="miguel fernández <miguel@hasura.io>"

ARG TARGETARCH

WORKDIR /app
COPY dist/hiring_log_gen_linux_$TARGETARCH/hiring_log_gen /

ENTRYPOINT ["/hiring_log_gen"]
