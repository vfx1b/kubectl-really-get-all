# kubectl-really-get-all plugin

Install instructions:

```
$ git clone https://github.com/vfx1b/kubectl-really-get-all.git

$ cd kubectl-really-get-all
$ go install
```
or
```
make install
```


Sample command:

```
$ kubectl really get all -n cert-manager
NAME				DATA	AGE
configmap/cert-manager-webhook	0	5h6m
configmap/kube-root-ca.crt	1	5h6m

NAME				ENDPOINTS	AGE
endpoint/cert-manager		10.244.0.7:9402	5h6m
endpoint/cert-manager-webhook	10.244.0.6:102505h6m

NAME						READY	STATUS	RESTARTSAGE
pod/cert-manager-59bf757d77-bsvb6		1/1	Running	0	5h7m
pod/cert-manager-cainjector-547c9b8f95-q9zsx	1/1	Running	0	5h7m
pod/cert-manager-webhook-6787f645b9-6blb7	1/1	Running	0	5h7m

NAME						TYPE			DATA	AGE
secret/cert-manager-webhook-ca			Opaque			3	5h6m
secret/sh.helm.release.v1.cert-manager.v1	helm.sh/release.v1	1	5h7m

NAME					SECRETS	AGE
serviceaccount/cert-manager		0	5h7m
serviceaccount/cert-manager-cainjector	0	5h7m
serviceaccount/cert-manager-webhook	0	5h7m
serviceaccount/default			0	5h7m

NAME				TYPE		CLUSTER-IP	EXTERNAL-IP	PORT(S)	AGE
service/cert-manager		ClusterIP	10.96.74.250	<none>		9402/TCP5h7m
service/cert-manager-webhook	ClusterIP	10.96.132.137	<none>		443/TCP	5h7m

NAME					READY	UP-TO-DATE	AVAILABLE	AGE
deployment/cert-manager			1/1	1		1		5h7m
deployment/cert-manager-cainjector	1/1	1		1		5h7m
deployment/cert-manager-webhook		1/1	1		1		5h7m

NAME						DESIRED	CURRENT	READY	AGE
replicaset/cert-manager-59bf757d77		1	1	1	5h7m
replicaset/cert-manager-cainjector-547c9b8f95	1	1	1	5h7m
replicaset/cert-manager-webhook-6787f645b9	1	1	1	5h7m

W0327 19:02:35.357883   55152 warnings.go:70] autoscaling/v2beta2 HorizontalPodAutoscaler is deprecated in v1.23+, unavailable in v1.26+; use autoscaling/v2 HorizontalPodAutoscaler
NAME						ROLE						AGE
rolebinding/cert-manager-webhook:dynamic-servingRole/cert-manager-webhook:dynamic-serving	5h7m

NAME						CREATED AT
role/cert-manager-webhook:dynamic-serving	2023-03-27T11:55:31Z

W0327 19:02:38.748372   55152 warnings.go:70] storage.k8s.io/v1beta1 CSIStorageCapacity is deprecated in v1.24+, unavailable in v1.27+; use storage.k8s.io/v1 CSIStorageCapacity
NAME					ADDRESSTYPE	PORTS	ENDPOINTS	AGE
endpointslice/cert-manager-trr46	IPv4		9402	10.244.0.7	5h7m
endpointslice/cert-manager-webhook-t47wgIPv4		10250	10.244.0.6	5h7m
```
Standard `kubectl get all` for comparison:

```
$ kubectl get all -n cert-manager
NAME                                           READY   STATUS    RESTARTS   AGE
pod/cert-manager-59bf757d77-bsvb6              1/1     Running   0          5h7m
pod/cert-manager-cainjector-547c9b8f95-q9zsx   1/1     Running   0          5h7m
pod/cert-manager-webhook-6787f645b9-6blb7      1/1     Running   0          5h7m

NAME                           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
service/cert-manager           ClusterIP   10.96.74.250    <none>        9402/TCP   5h7m
service/cert-manager-webhook   ClusterIP   10.96.132.137   <none>        443/TCP    5h7m

NAME                                      READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/cert-manager              1/1     1            1           5h7m
deployment.apps/cert-manager-cainjector   1/1     1            1           5h7m
deployment.apps/cert-manager-webhook      1/1     1            1           5h7m

NAME                                                 DESIRED   CURRENT   READY   AGE
replicaset.apps/cert-manager-59bf757d77              1         1         1       5h7m
replicaset.apps/cert-manager-cainjector-547c9b8f95   1         1         1       5h7m
replicaset.apps/cert-manager-webhook-6787f645b9      1         1         1       5h7m
```

TODO:
* handle kube env
* better parameters