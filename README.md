# Demo HTTP Service - SLO driven Monitoring


## Service Definition:
A stateless HTTP API that serves user requests over HTTP.  

Users care about how many requests are successful and what is the latency. These two will be made as service level indicators.  

## SLIs (Service Level Indicators)

### 1. Request Success Rate

**Question:** What percentage of user requests succeed?  

**Metric Source:**
Total requests  
Total errors  

**PromQL:**  
```
1 - (
  sum(rate(demo_http_errors_total[5m]))
  /
  sum(rate(demo_http_request_duration_seconds_count[5m]))
)
```

This yields a ratio between 0 and 1.  

### 2. Request Latency (P95)

**Question:** How slow is the slowest 5% of user requests?  

**PromQL:**
```
histogram_quantile(
  0.95,
  sum(rate(demo_http_request_duration_seconds_bucket[5m])) by (le)
)
```


## SLOs (Service Level Objectives)

1. Availability: 99% success rate over 5 minutes  
2. Latency: P95<=1 sec  

Short spikes are tolerated. Alerts fire only on sustained violations.  


## Alerting Philosophy  

Alerts are based on SLO violations rather than raw metrics.  
The goal is to page only when the users are experiencing some issue/pain.  

## Monitoring Design Decisions

1. 99% availability chosen to allow small error budget.
2. 5m window balances sensitivity and noise.  
3. P95 latency reflects tail-user experience. 




