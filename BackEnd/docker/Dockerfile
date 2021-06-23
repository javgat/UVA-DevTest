FROM debian:buster-slim

RUN apt update && apt install -y g++

COPY run.sh /bin/run_eval

RUN chmod +x /bin/run_eval

ENTRYPOINT /bin/run_eval