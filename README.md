# kubectl-really-get-all plugin

Install instructions:

```$ go install github.com/vfx1b/kubectl-really-get-all```

Sample command:

```

$ kubectl really get all -n cert-manager
NAME			DATA	AGE
cert-manager-webhook	0	85m
kube-root-ca.crt	1	85m

NAME			ENDPOINTS	AGE
cert-manager		10.244.0.7:9402	85m
cert-manager-webhook	10.244.0.6:1025085m

NAME					READY	STATUS	RESTARTSAGE
cert-manager-59bf757d77-bsvb6		1/1	Running	0	85m	10.244.0.7	kind-control-plane	<none>	<none>
cert-manager-cainjector-547c9b8f95-q9zsx1/1	Running	0	85m	10.244.0.5	kind-control-plane	<none>	<none>
cert-manager-webhook-6787f645b9-6blb7	1/1	Running	0	85m	10.244.0.6	kind-control-plane	<none>	<none>

NAME					TYPE			DATA	AGE
cert-manager-webhook-ca			Opaque			3	85m
sh.helm.release.v1.cert-manager.v1	helm.sh/release.v1	1	85m

NAME			SECRETS	AGE
cert-manager		0	85m
cert-manager-cainjector	0	85m
cert-manager-webhook	0	85m
default			0	85m

NAME			TYPE		CLUSTER-IP	EXTERNAL-IP	PORT(S)	AGE
cert-manager		ClusterIP	10.96.74.250	<none>		9402/TCP85m	app.kubernetes.io/component=controller,app.kubernetes.io/instance=cert-manager,app.kubernetes.io/name=cert-manager
cert-manager-webhook	ClusterIP	10.96.132.137	<none>		443/TCP	85m	app.kubernetes.io/component=webhook,app.kubernetes.io/instance=cert-manager,app.kubernetes.io/name=webhook

NAME			READY	UP-TO-DATE	AVAILABLE	AGE
cert-manager		1/1	1		1		85m	cert-manager-controller	quay.io/jetstack/cert-manager-controller:v1.11.0app.kubernetes.io/component=controller,app.kubernetes.io/instance=cert-manager,app.kubernetes.io/name=cert-manager
cert-manager-cainjector	1/1	1		1		85m	cert-manager-cainjector	quay.io/jetstack/cert-manager-cainjector:v1.11.0app.kubernetes.io/component=cainjector,app.kubernetes.io/instance=cert-manager,app.kubernetes.io/name=cainjector
cert-manager-webhook	1/1	1		1		85m	cert-manager-webhook	quay.io/jetstack/cert-manager-webhook:v1.11.0	app.kubernetes.io/component=webhook,app.kubernetes.io/instance=cert-manager,app.kubernetes.io/name=webhook

NAME					DESIRED	CURRENT	READY	AGE
cert-manager-59bf757d77			1	1	1	85m	cert-manager-controller	quay.io/jetstack/cert-manager-controller:v1.11.0app.kubernetes.io/component=controller,app.kubernetes.io/instance=cert-manager,app.kubernetes.io/name=cert-manager,pod-template-hash=59bf757d77
cert-manager-cainjector-547c9b8f95	1	1	1	85m	cert-manager-cainjector	quay.io/jetstack/cert-manager-cainjector:v1.11.0app.kubernetes.io/component=cainjector,app.kubernetes.io/instance=cert-manager,app.kubernetes.io/name=cainjector,pod-template-hash=547c9b8f95
cert-manager-webhook-6787f645b9		1	1	1	85m	cert-manager-webhook	quay.io/jetstack/cert-manager-webhook:v1.11.0	app.kubernetes.io/component=webhook,app.kubernetes.io/instance=cert-manager,app.kubernetes.io/name=webhook,pod-template-hash=6787f645b9

W0327 15:20:49.756478   26559 warnings.go:70] autoscaling/v2beta2 HorizontalPodAutoscaler is deprecated in v1.23+, unavailable in v1.26+; use autoscaling/v2 HorizontalPodAutoscaler
NAME					ROLE						AGE
cert-manager-webhook:dynamic-serving	Role/cert-manager-webhook:dynamic-serving	85m	cert-manager/cert-manager-webhook

NAME					CREATED AT
cert-manager-webhook:dynamic-serving	2023-03-27T11:55:31Z

W0327 15:20:53.159704   26559 warnings.go:70] storage.k8s.io/v1beta1 CSIStorageCapacity is deprecated in v1.24+, unavailable in v1.27+; use storage.k8s.io/v1 CSIStorageCapacity
NAME				ADDRESSTYPE	PORTS	ENDPOINTS	AGE
cert-manager-trr46		IPv4		9402	10.244.0.7	85m
cert-manager-webhook-t47wg	IPv4		10250	10.244.0.6	85m

```