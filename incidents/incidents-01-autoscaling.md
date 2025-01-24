# Incident: CPU Saturation and Autoscaling Event

## Summary
Sustained high-latency traffic caused increased CPU Utilization in the demo-service triggering horizontal pod autoscaling.  

## Impact
Users experienced elevated P95 latency before autoscalingstabilized the service. No request failures occurred.

## Detection
Detected via Grafana dashboards showing increased CPU usage and rising P95 latency. HPA events confirmed scaling activity.  

## Root Cause
Increased request processing time from /slow endpoint led to CPU saturation on existing pods, exceeding HPA target utilization.  

## Resolution
The Horizontal Pod Autoscaler increased the number of replicas, distributing load across more podsand reducing CPU pressure.

## Lessons Learned
Autoscaling effectively mitigates CPU-bound latency issues when resource requests and metrics are correctly configured. Observability dashboards are essential to correlate scaling behavior with user-facing latency.
