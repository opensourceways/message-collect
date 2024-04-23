import json
from dataclasses import dataclass
from fedora_messaging import api, config
from kafka import KafkaProducer
from typing import Dict


@dataclass
class Event:
    Key: str
    Header: Dict[str, str]
    Body: object


# 创建生产者
producer = KafkaProducer(bootstrap_servers='localhost:9092', api_version=(2, 8, 1))

# Kafka 主题
topic = 'eur_build_raw'

config.conf.setup_logging()


def json_marshal(obj):
    # 将对象编码为 JSON 格式的字符串
    json_string = json.dumps(
        obj,
        default=lambda o: o.__dict__,
        sort_keys=True,
        indent=4)
    # 将字符串转换为字节数组
    json_bytes = json_string.encode('utf-8')
    return json_bytes


def print_eur_message(message):
    # all eur related messages https://apps.fedoraproject.org/datagrepper/raw?category=copr
    if message.topic == "org.fedoraproject.prod.copr.build.end" or message.topic == "org.fedoraproject.prod.copr.build.start":
        print("******** Received eur event: ********\n")
        print(message)

        producer.send(topic, json_marshal(message))
    else:
        print("******** Message is not triggered by eur, skipped ********")


def main():
    api.consume(callback=print_eur_message)


main()
