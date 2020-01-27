# none: the number of available CPUs 1 is less than the required
Solution: 
```
sudo minikube start --vm-driver=none --extra-config=kubeadm.ignore-preflight-errors=NumCPU --force
```
Reference: https://github.com/kubernetes/minikube/issues/5010

