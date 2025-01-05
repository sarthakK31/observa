# Observa
project for grouping and testing out the monitoring stack consisting of prometheus, grafana and alert manager. Objective is to see them in action on a mock kubernetes cluster deployed via kind.   
kind allows us to simulate real life kubeadm closely with the worker node and control plane structure among other benefits.  

so we also need to test the system under load and various alert metrics hence a small go server can come in handy for testing a bunch of parameters.
