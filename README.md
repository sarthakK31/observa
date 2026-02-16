# Observa — Kubernetes Observability & Reliability Stack

A production-style observability and reliability platform built on a multi-node Kubernetes cluster to simulate real-world SRE workflows including:

- SLO-driven monitoring  
- Metrics collection & visualization  
- Alerting based on user impact  
- Autoscaling under load  
- Centralized logging  
- Metrics-to-logs correlation  
- Incident simulation & root cause analysis  

---

# 1. Service Definition – Demo HTTP Service

At the core of the platform is a **stateless HTTP API** deployed on Kubernetes.

From a user perspective, only two things matter:

1. Are requests successful?
2. How fast are responses?

These concerns are formalized as:

- **SLIs (Service Level Indicators)**
- **SLOs (Service Level Objectives)**

The system is designed so alerts reflect **real user impact**, not just metric noise.

---

# 2. Architecture Overview

## High-Level Flow

```
User Load
   ↓
Kubernetes Service
   ↓
Demo Microservice (Replicas)
   ↓
Metrics → Prometheus → Grafana
Logs → Promtail → Loki → Grafana
CPU Metrics → HPA → Pod Scaling
Alerts → Alertmanager
```

## Core Stack Components

- Kubernetes (kind multi-node cluster)
- Prometheus (kube-prometheus-stack)
- Grafana
- Alertmanager
- Loki (multi-tenant mode)
- Promtail (DaemonSet log shipping)
- Horizontal Pod Autoscaler (HPA)

---

# 3. SLO-Driven Monitoring Model

## SLIs (Service Level Indicators)

### 🔹 1. Request Success Rate (Availability)

**User Question:**  
What percentage of user requests succeed?

**Metric Sources:**

- `demo_http_request_duration_seconds_count` (total requests)
- `demo_http_errors_total` (total errors)

**PromQL:**

```promql
1 - (
  sum(rate(demo_http_errors_total[5m]))
  /
  sum(rate(demo_http_request_duration_seconds_count[5m]))
)
```

Returns a ratio between **0 and 1** over a rolling 5-minute window.

---

### 🔹 2. Request Latency (P95)

**User Question:**  
How slow are the slowest 5% of requests?

We measure tail latency using histogram quantiles.

**PromQL:**

```promql
histogram_quantile(
  0.95,
  sum(rate(demo_http_request_duration_seconds_bucket[5m])) by (le)
)
```

P95 reflects real user experience better than averages.

---

## SLOs (Service Level Objectives)

1. **Availability:** ≥ 99% success rate (5-minute window)  
2. **Latency:** P95 ≤ 1 second  

Short spikes are tolerated. Alerts fire only on sustained violations using `for` conditions.

---

# 4. Observability Components

## Metrics – Prometheus

Collected and visualized:

- Request rate (R)
- Error rate (E)
- Duration (D / P95)
- CPU utilization
- Per-pod breakdown

Example:

```promql
sum(rate(demo_http_request_duration_seconds_count[5m])) by (pod)
```

The service follows the **RED methodology**:
- Rate
- Errors
- Duration

---

## Alerts – SLO-Based Alerting

Alerts are defined via `PrometheusRule` CRDs.

### Philosophy

Alerts are based on **SLO violations**, not raw metric thresholds.

Goal:
- Page only when users are impacted
- Avoid alert fatigue
- Tie alerts directly to user experience

Example latency alert:

```promql
histogram_quantile(
  0.95,
  sum(rate(demo_http_request_duration_seconds_bucket[5m])) by (le)
) > 0.5
```

Sustained breach detection uses `for:` conditions.

---

## Autoscaling – Horizontal Pod Autoscaler

- CPU-based scaling
- Dynamic replica adjustments under load
- Validated scale-up and scale-down behavior

Observed via:

```bash
watch -n 2 kubectl get hpa
```

Scaling behavior is validated during simulated incidents.

---

## Centralized Logging – Loki + Promtail

- Promtail deployed as DaemonSet
- Node-local container log scraping
- containerd-compatible configuration
- Multi-tenant Loki (`tenant_id` / `X-Scope-OrgID`)
- LogQL filtering & aggregation

Example:

```logql
{namespace="default",pod="demo-service-xxxx"} |= "slow"
```

---

# 5. Metrics <-> Logs Correlation

Grafana configured with clickable drilldowns:

- Metrics grouped by `pod`
- Data links use `${__field.labels.pod}`
- Clicking a metric opens logs in Loki for that pod

Workflow:

Metric spike → Identify pod → Inspect logs → Root cause

This enables realistic SRE debugging flows.

---

# 6. Incident Simulation & Reliability Validation

To validate alerting, scaling, and SLO behavior, the following scenarios were simulated:

- Traffic spike
- Artificial latency injection
- Error injection
- Sustained SLO violations
- HPA scale-up validation
- HPA scale-down validation
- Alert firing validation
- Log-based root cause analysis

### Demo Flow

1. Healthy baseline
2. Generate load
3. Latency spike
4. Alert fires
5. HPA scales
6. Click metric → View logs
7. Root cause identified

Each incident is documented with:

- Detection timeline
- Root cause
- Impact analysis
- Lessons learned

---

# 7. Monitoring Design Decisions

1. **99% Availability**
   - Allows controlled error budget
   - Realistic production trade-off

2. **5-Minute Evaluation Window**
   - Balances sensitivity and noise

3. **P95 Latency**
   - Captures tail-user experience

4. **Multi-Node kind Cluster**
   - Realistic scheduling behavior

5. **DaemonSet Promtail**
   - Node-local log visibility

6. **Multi-Tenant Loki**
   - Production-style configuration complexity

---

# 8. Key Debugging Challenges Solved

- containerd log path mismatch in kind
- Promtail file descriptor exhaustion
- Loki multi-tenant ingestion failures
- Scrape relabel misconfiguration
- Promtail positions file offset behavior
- Grafana data-link interpolation issues
- DaemonSet log visibility constraints

---

# 9. Cleanup

```bash
kind delete cluster --name observa
docker system prune -a
```

---

# 10. Lessons Learned

- Discovery != Tailing in log pipelines
- SLO-based alerts reduce noise significantly
- Multi-tenant Loki requires HTTP tenant headers
- containerd log paths differ from Docker
- Promtail readiness probes can mislead in low-volume systems
- Observability is about correlating signals, not just collecting them
- Autoscaling must be validated during real load events

---

# What This Project Demonstrates

- Production-style observability architecture
- SLO-driven monitoring design
- Alerting aligned with user impact
- Autoscaling validation under stress
- Centralized multi-tenant logging
- Metric-to-log drilldowns
- Realistic incident response workflow
- Kubernetes-native reliability engineering

---

# Future Improvements

- Add distributed tracing (Grafana Tempo)
- Implement SLO burn-rate alerts
- Add recording rules
- Deploy via Terraform (full IaC)
- Introduce error-budget reporting dashboards

---

**Built to simulate real SRE incident response workflows on Kubernetes using SLO-driven observability principles.**

**Please view the following link for checking the project-screenshots**
** **

### **Here is a video demo of the project in action**
<video width="640" height="480" controls>
  <source src="project-screenshots/workflow-demo.webm" type="video/mp4">
  Your browser does not support the video tag.
</video>
