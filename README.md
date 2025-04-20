# FlowWatch

A cloud-native platform that pulls public river current data, exposes it via a REST API, and visualizes it on a simple frontend. The app is deployed on Kubernetes and includes all the SRE/DevOps best practices.

## 🔧 Tech Stack


| Layer | Tech |
|---|---|
| Data Source | USGS Water Services API (https://waterservices.usgs.gov/) |
| Backend | Go or Python FastAPI app to fetch & process data |
| Frontend | React or static HTML/JS with charts (Chart.js, D3.js) |
| API Schedule | CronJob or Sidekiq-style worker in Kubernetes |
| Deployment | Kubernetes (Minikube, kind, GKE, or EKS) |
| IaC | Terraform (for cloud infra) + Helm for K8s |
| Monitoring | Prometheus + Grafana |
| CI/CD | GitHub Actions |
| Storage | PostgreSQL, Redis, or just ephemeral storage |


## 📦 Core Features

1. Data Collection

- Periodically fetch river current data (flow rate, gage height) from public APIs like USGS.
- Store in memory or lightweight DB (Redis/Postgres).
- Cache responses for frontend use.

2. REST API

- Expose latest river data via endpoints like:
- /api/river/current?site=09380000
- /api/river/history?days=7

3. Visualization

- Simple React or static dashboard showing:
- Line charts for river flow rate over time
- Location marker maps using Leaflet or Mapbox

4. Kubernetes Hosting

- Deploy API + frontend on Kubernetes
- Use CronJobs or sidecar jobs to collect data
- Add resource requests/limits, liveness probes, HPA

5. Monitoring & Alerts

- Prometheus tracks:
- API uptime
- Data fetch success rate
- Grafana dashboards for:
    - Request latency
    - River data values
    - Alertmanager for failures (email/webhook)

6. CI/CD

- GitHub Actions builds + lints code
- Auto-deploy to K8s cluster on main branch merge