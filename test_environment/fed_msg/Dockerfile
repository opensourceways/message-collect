FROM python:3.10
LABEL authors="tommylike"

RUN pip install fedora-messaging  && pip install kafka-python && apt update

COPY config.toml /root/config.toml
COPY main.py /root/main.py
COPY certs/cacert.pem /etc/fedora-messaging/cacert.pem
COPY certs/fedora-cert.pem /etc/fedora-messaging/fedora-cert.pem
COPY certs/fedora-key.pem /etc/fedora-messaging/fedora-key.pem

ENV FEDORA_MESSAGING_CONF /root/config.toml

ENTRYPOINT ["python", "/root/main.py"]