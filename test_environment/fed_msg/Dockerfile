FROM python:3.10
LABEL authors="tommylike"

RUN pip install fedora-messaging && apt update && mkdir /etc/fedora-messaging/
RUN curl https://raw.githubusercontent.com/fedora-infra/fedora-messaging/stable/configs/fedora-key.pem -o /etc/fedora-messaging/fedora-key.pem && \
    curl https://raw.githubusercontent.com/fedora-infra/fedora-messaging/stable/configs/fedora-cert.pem -o /etc/fedora-messaging/fedora-cert.pem && \
    curl https://raw.githubusercontent.com/fedora-infra/fedora-messaging/stable/configs/cacert.pem -o /etc/fedora-messaging/cacert.pem

COPY config.toml /root/config.toml
COPY main.py /root/main.py

ENV FEDORA_MESSAGING_CONF /root/config.toml

ENTRYPOINT ["python", "/root/main.py"]