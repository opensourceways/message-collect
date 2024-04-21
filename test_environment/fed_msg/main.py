from fedora_messaging import api, config

config.conf.setup_logging()


def print_eur_message(message):
    # all eur related messages https://apps.fedoraproject.org/datagrepper/raw?category=copr
    if message.topic == "org.fedoraproject.prod.copr.build.end" or message.topic == "org.fedoraproject.prod.copr.build.start":
        print("******** Received eur event: ********\n")
        print(message)
    else:
        print("******** Message is not triggered by eur, skipped ********")


def main():
    api.consume(callback=print_eur_message)

main()