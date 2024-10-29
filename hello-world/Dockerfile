FROM ubuntu:latest

RUN apt update -y && apt upgrade -y && \
	apt install -y \
	gcc \
	strace


WORKDIR /app

COPY ./entrypoint.sh ./main.c ./

RUN gcc -o ./main ./main.c

RUN chmod u+x ./entrypoint.sh

CMD ["./entrypoint.sh"]
