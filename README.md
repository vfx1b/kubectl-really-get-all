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
configmaps/cert-manager-webhook	0	133m
configmaps/kube-root-ca.crt	1	133m

NAME				ENDPOINTS	AGE
endpoints/cert-manager		10.244.0.7:9402	133m
endpoints/cert-manager-webhook	10.244.0.6:10250133m

NAME						READY	STATUS	RESTARTSAGE
pods/cert-manager-59bf757d77-bsvb6		1/1	Running	0	133m
pods/cert-manager-cainjector-547c9b8f95-q9zsx	1/1	Running	0	133m
pods/cert-manager-webhook-6787f645b9-6blb7	1/1	Running	0	133m

NAME						TYPE			DATA	AGE
secrets/cert-manager-webhook-ca			Opaque			3	133m
secrets/sh.helm.release.v1.cert-manager.v1	helm.sh/release.v1	1	133m

NAME					SECRETS	AGE
serviceaccounts/cert-manager		0	133m
serviceaccounts/cert-manager-cainjector	0	133m
serviceaccounts/cert-manager-webhook	0	133m
serviceaccounts/default			0	133m

NAME				TYPE		CLUSTER-IP	EXTERNAL-IP	PORT(S)	AGE
services/cert-manager		ClusterIP	10.96.74.250	<none>		9402/TCP133m
services/cert-manager-webhook	ClusterIP	10.96.132.137	<none>		443/TCP	133m

NAME					READY	UP-TO-DATE	AVAILABLE	AGE
deployments/cert-manager		1/1	1		1		133m
deployments/cert-manager-cainjector	1/1	1		1		133m
deployments/cert-manager-webhook	1/1	1		1		133m

NAME						DESIRED	CURRENT	READY	AGE
replicasets/cert-manager-59bf757d77		1	1	1	133m
replicasets/cert-manager-cainjector-547c9b8f95	1	1	1	133m
replicasets/cert-manager-webhook-6787f645b9	1	1	1	133m

W0327 16:08:48.040441   33393 warnings.go:70] autoscaling/v2beta2 HorizontalPodAutoscaler is deprecated in v1.23+, unavailable in v1.26+; use autoscaling/v2 HorizontalPodAutoscaler
NAME							ROLE						AGE
rolebindings/cert-manager-webhook:dynamic-serving	Role/cert-manager-webhook:dynamic-serving	133m

NAME						CREATED AT
roles/cert-manager-webhook:dynamic-serving	2023-03-27T11:55:31Z

W0327 16:08:51.452319   33393 warnings.go:70] storage.k8s.io/v1beta1 CSIStorageCapacity is deprecated in v1.24+, unavailable in v1.27+; use storage.k8s.io/v1 CSIStorageCapacity
NAME						ADDRESSTYPE	PORTS	ENDPOINTS	AGE
endpointslices/cert-manager-trr46		IPv4		9402	10.244.0.7	133m
endpointslices/cert-manager-webhook-t47wg	IPv4		10250	10.244.0.6	133m



```

TODO:
* handle kube env
* better parameters