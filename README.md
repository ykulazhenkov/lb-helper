###Usage

```
$ kubectl apply -f service.yaml -f replicaSet.yaml

$ kubectl get po
NAME                READY   STATUS    RESTARTS   AGE
lbhelper-rs-4wkn9   1/1     Running   0          5s
lbhelper-rs-k4msz   1/1     Running   0          5s
lbhelper-rs-rnvm5   1/1     Running   0          5s

$ kubectl get service
  NAME         TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
  kubernetes   ClusterIP      10.96.0.1       <none>        443/TCP          23h
  lbhelperlb   LoadBalancer   10.109.230.20   172.100.0.1   8080:30317/TCP   16s


# from client node

$ curl http://172.100.0.1:8080
{
   "node_name":"node1",
   "node_ip":"10.249.233.26",
   "pod_name":"lbhelper-rs-9vg2h",
   "pod_namespace":"default",
   "pod_ip":"10.244.0.14",
   "client_ip":"192.168.30.2:50076",
   "call_count":1
}
```
