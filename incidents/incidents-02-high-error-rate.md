# Incident: Elevated HTTP Error Rate

## Summary
A sustained increase in error responses triggered the high error rate alert for demo-service.  

## Impact
A significant number of user requests returned HTTP 5xx errors during the incident window.

## Detection
The Prometheus alert 'DemoServiceHighErrorRate' fired after the error rate exceeded the defined SLO threshold for more than 2 minutes.  

## Root Cause
Intentional traffic to the /error generated repeated server-side errors, increasing the overall error ratio.

## Resolution
Error-inducing traffic was stopped, allowing the error rate to fall belowthe alert threshold and the alert to resolve automatically.  

## Lessons Learned
Error-rate-based alerts accurately capture user-visible failures. Alert thresholds and evaluation windows prevented transient errors from causing unnecessary pages.    
