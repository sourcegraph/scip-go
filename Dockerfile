FROM golang:1.20.3@sha256:bcc311ec9655c350df3899611fdf134806f97a3e3b2c06c2b5c0696428503814

RUN curl -L https://sourcegraph.com/.api/src-cli/src_linux_amd64 -o /usr/bin/src && chmod +x /usr/bin/src

COPY scip-go /usr/bin/
