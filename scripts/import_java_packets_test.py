#!/usr/bin/env python3
from scripts.import_java_packets import import_packets_wiki


def test_packet_counts():
    packets = import_packets_wiki()

    assert len(packets["handshaking"]["clientbound"]) == 0
    assert len(packets["handshaking"]["serverbound"]) == 1

    assert len(packets["status"]["clientbound"]) == 2
    assert len(packets["status"]["serverbound"]) == 2

    assert len(packets["login"]["clientbound"]) == 6


if __name__ == "__main__":
    test_packet_counts()
    print("OK")
