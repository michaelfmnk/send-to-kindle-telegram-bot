# TODO: refactor dockerfile
FROM amd64/ubuntu
COPY --from=golang:1.16.11-bullseye /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

RUN go env -w GOPROXY=direct GOFLAGS="-insecure"
ENV TZ=Europe/Minsk
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git && \
    apt-get install wget && \
    apt-get install -y python
RUN apt-get install -y ffmpeg libsm6 libxext6
RUN wget -nv -O- https://download.calibre-ebook.com/linux-installer.sh | sh /dev/stdin
RUN mkdir -p files
WORKDIR /app

COPY . .

RUN go build

RUN chmod +x ./send-to-kindle-telegram-bot
CMD ["./send-to-kindle-telegram-bot"]