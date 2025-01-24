# Incident: Latency Spikes without Alert Trigger

## Summary
Intermittent latency spikes were visible in dashboards but did not trigger a latency alert.  

## Impact
Some requests experienced higher latency, but overall service performance remained within the defined SLO.  

## Detection
Observed manually via Grafana P95 latency dashboards using short evaulation windows.  

## Root Cause
Latency spikes were transient and did not violate the sustained SLO conditions defined in the Prometheus alert (5m window with hold time)  

## Resolution
No action was required. The alert correctly remained inactive as the service level objectives were not breached.  

## Lessons Learned
Dashboards and alerts serve different purposes. Alerts should be conservative to avoid noise, while dashboards provide high-resolution visibility into transient behaviour.