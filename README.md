# 🛍️ Simple E-Commerce — Microservices Architecture

This project is a simple e-commerce web application built to explore and implement **Microservices Architecture**.
It serves as a playground to understand distributed systems, where users can browse products, manage carts, and simulate checkout through independent, decoupled services.

---

## 🚀 Overview

The main goal of this project is to learn how to build a scalable and resilient backend system using modern technologies. The application is divided into smaller independent services written in **Go (Gin framework)**.

Key features of this architecture:
- **Decoupled Services:** Each domain (Auth, Product, Order) has its own service.
- **Independent Databases:** Each service manages its own database to ensure loose coupling.
- **Event-Driven Communication:** Services communicate asynchronously using **Kafka** to handle complex workflows like order processing.

This setup helps in exploring concepts like scalability, fault tolerance, and service communication patterns in a distributed environment.

---

## ⚙️ Tech Stack

### 🖥️ Frontend
- **React + TypeScript**
- Axios for API communication
- Basic state management (React Context or simple hooks)
- Styled Components / TailwindCSS for UI styling

### 🔗 Backend Services
- **Go + Gin Framework**
- **GORM ORM**
- **Kafka** for asynchronous messaging and event handling
- **Docker** for containerization
- RESTful APIs for client-service communication

### 🗄️ Database
- **PostgreSQL** (Each service has its own isolated database)

---

## 🧩 Architecture Breakdown

The system is composed of the following services:

- **Auth Service** – Manages user registration, login, and JWT authentication.
- **Product Service** – Handles the product catalog, pricing, and inventory management.
- **Order Service** – Processes orders and manages checkout workflows.
- **Kafka Broker** – Facilitates event-driven communication (e.g., publishing `OrderCreated` events to trigger inventory updates).
- **API Gateway (Optional)** – Can be added for routing and request aggregation.

---

## 🧠 Learning Outcomes

- Designing and building a **microservices** architecture from scratch.
- Managing **data consistency** in a distributed system.
- Implementing **Kafka** for event-driven messaging between Go services.
- Using **Docker** and Docker Compose to orchestrate multiple services and databases.
- Understanding the complexities of distributed transactions and service independence.

---

## 💡 Future Improvements

- Add real payment gateway integration.
- Implement distributed tracing and logging (Jaeger / Prometheus).
- Improve frontend UI and UX.
- Add CI/CD pipeline for automated deployment.
- Deploy the cluster to a cloud provider (e.g., AWS, GCP, or DigitalOcean).

---

## 📦 How to Run

### 🖥️ Frontend
```bash
cd frontend
npm install
npm run dev
```


### ⚙️ Backend
```bash
cd backend-microservices
docker-compose up --build
```

---

## 🧑‍💻 Author
Choirul Anam
Computer Science student exploring full-stack development, backend architecture, and distributed systems.

---

## 📄 License
This project is open source and free to use for learning and educational purposes.


