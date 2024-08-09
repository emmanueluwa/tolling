# Toll Calculator

# Microservices Project: GPS Tracking and Invoicing System

## Project Overview

This project is a microservices-based system designed to track the movements of trucks on the road using GPS signals transmitted from an On-Board Unit (OBU). The system is composed of several services that work together to collect GPS data, calculate distances traveled, and generate invoices for the trucking companies. Below is a brief summary of each service within the project:

### 1. **Receiver Service**
- **Functionality:** Receives GPS coordinates from the OBUs installed in the trucks. 
- **Process:** The GPS data is collected at regular intervals and sent to the Receiver service.
- **Output:** The received coordinates are placed onto a Kafka queue for further processing.

### 2. **Distance Calculator Service**
- **Functionality:** Retrieves GPS coordinates from the Kafka queue and calculates the distance traveled by each truck.
- **Process:** This service listens to the Kafka queue, processes the coordinates, and computes the distance between the received points.
- **Output:** The calculated distance is passed on to the next service for invoicing.

### 3. **Invoicer Service**
- **Functionality:** Generates invoices based on the calculated distances.
- **Process:** This service uses an Invoice Calculator to determine the cost based on the distance traveled. The generated invoices are then stored in a database for record-keeping.
- **Output:** Invoices are made available through an API gateway for interaction with a frontend application.

### 4. **API Gateway**
- **Functionality:** Serves as the interface for external communication, providing endpoints for accessing the various services.
- **Process:** The API gateway aggregates the services, enabling clients to interact with the system via a unified interface. This is tested using Postman.
- **Output:** Facilitates the interaction between the frontend and backend services.

## Technology Stack
- **Programming Language:** Golang
- **Message Queue:** Kafka
- **Containerization and Orchestration:** Docker (used for running Kafka and other services)
- **API Testing Tool:** Postman

## Usage
The system can be tested and interacted with via the API Gateway. For testing and development purposes, Postman is used to simulate client requests and validate the functionality of each service.




---

docker run --name kafka \
  -p 9092:9092 \
    -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181 \
      -e ALLOW_PLAINTEXT_LISTENER=yes \
        -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true \
          -e KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT \
            bitnami/kafka:latest

---


