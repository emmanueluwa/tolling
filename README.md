# Toll Calculator

---

docker run --name kafka \
  -p 9092:9092 \
    -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181 \
      -e ALLOW_PLAINTEXT_LISTENER=yes \
        -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true \
          -e KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT \
            bitnami/kafka:latest

---


