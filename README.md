# Rlist 

Small command util written in Go for extract data from a K8S cluster

## Usage

```
./rlist -h
Usage of ./rlist:
  -kubeconfig
        show the kubeconfig loaded
  -n string
        the namespace to show
  -nodes
        show the nodes
```

```
./rlist -n kube-system -nodes
2020/11/02 04:46:00 There are 1 nodes in the cluster
2020/11/02 04:46:00 There are 7 pods in the kube-system namespace
```