# OTP Login & Document Verification Platform

A small multi-service system written in Go, with a React frontend, simulating OTP-based user authentication and fake document verification.

This project is a learning exercise to gain experience with Go web development, modular service design, Redis/Postgres integration, and future extensibility.

---

## 🎯 Goals

- Build a minimal but extensible full-stack system using Go (backend) and React (frontend)
- Use Gin for web APIs, GORM for database interactions, Redis for OTP caching
- Simulate a realistic login and document workflow with:
  - OTP authentication
  - Document upload
  - Deterministic document verification based on hashing
- Practice proper Go project layout and microservice boundaries
- Use Docker Compose to manage infrastructure (Redis, Postgres, etc.)
- Explore design patterns that enable future features (queues, workers, polling, etc.)
- Keep the frontend simple using React and Bootstrap

---

## 🧱 Services

| Name           | Description                                 |
| -------------- | ------------------------------------------- |
| `auth`         | Handles OTP generation and login logic      |
| `verification` | Handles document uploads and verification   |
| `frontend`     | Simple React UI for login and document flow |

---

## 🚧 Project Status

- ✅ Phase 1: PoC design and scaffolding
- ⏳ Phase 2: User persistence and token auth
- 🔜 Phase 3: Queue-based document verification
- 🔜 Phase 4: Monitoring and observability
- 🔜 Phase 5: Security and hardening

---

## 🐳 Stack

- Go 1.20+ with Gin and GORM
- PostgreSQL for persistent data
- Redis for OTP storage
- React + Bootstrap for UI
- Docker Compose for local development

---

## ⚙️ Getting Started

TBD — see `deployments/docker-compose.yml` for service setup (coming soon)
