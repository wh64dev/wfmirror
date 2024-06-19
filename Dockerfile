FROM golang:1.22.4-bookworm

WORKDIR /mirror

COPY . .
RUN apt update
RUN apt install make -y
RUN apt install sqlite3 -y

RUN ./configure
RUN make

EXPOSE 8080

ENTRYPOINT ["/mirror/wfmirror", "-S", "-C"]
