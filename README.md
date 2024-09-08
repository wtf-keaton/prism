# Prism - HR Platform 🌈👥

![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v2.40.1-00ADD8?style=for-the-badge&logo=go)
![Bootstrap](https://img.shields.io/badge/Bootstrap-v5.2-7952B3?style=for-the-badge&logo=bootstrap)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13+-4169E1?style=for-the-badge&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-20.10.21-2496ED?style=for-the-badge&logo=docker)

Prism is a modern HR platform designed to streamline the recruitment process, providing a seamless experience for both candidates and HR professionals. 🚀

## ✨ Features

- **For Candidates:**
  - 🔍 Browse available job openings
  - 📝 Apply for positions
  - 📊 Track application status

- **For HR Professionals:**
  - 📢 Create and manage job postings
  - 👀 Review candidate applications and resumes
  - 🔄 Manage the hiring pipeline with customizable stages
  - 🗄️ Archive job postings

## 🛠️ Technology Stack

- Backend: Go (Golang) 🐹
- UI Framework: Bootstrap 🎨
- Database: PostgreSQL 🐘
- Message Broker: Kafka 📬
- Caching: Redis 🗃️
- Monitoring: Prometheus & Grafana 📈

## 🏗️ Microservices Architecture

Prism is built on a microservices architecture, consisting of the following services:

1. 🔐 Auth Service
2. 👤 User Service
3. 💼 Vacancy Service
4. 📮 Api Gateway
5. 🔔 Notification Service
6. 🖥️ Frontend Service

## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- Docker and Docker Compose
- PostgreSQL 13+
- Kafka
- Redis

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/wtf-keaton/prism.git
   cd prism
   ```

2. Set up environment variables:
   ```
   cp .env.example .env
   ```
   Edit the `.env` file with your configuration.

3. Build and run all services (including frontend):
   ```
   docker-compose up --build
   ```

4. The services should now be running. You can access the frontend at `http://localhost:3000` and the API at `http://localhost:8080`.

## 💻 Development

To run the frontend service locally for development:

1. Navigate to the frontend service directory:
   ```
   cd services/frontend
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the service:
   ```
   go run cmd/main.go
   ```

The frontend will be available at `http://localhost:3000`.

## 📚 API Documentation

API documentation is available via Swagger UI. After starting the services, you can access it at:

`http://localhost:8080/swagger/index.html`

## 🤝 Contributing

We welcome contributions to Prism! Please see our [Contributing Guide](CONTRIBUTING.md) for more details.

## 🧪 Testing

To run the test suite:

```
go test ./...
```

## 📊 Monitoring

Prometheus metrics are exposed at `/metrics` endpoint for each service. Grafana dashboards are available to visualize these metrics.

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Contact

For any queries or support, please contact our team at support@libretto.store.

---

Made with ❤️ by the Prism team
