module github.com/ibrokethecloud/rancher-rbac-lister

go 1.14

replace (
	github.com/knative/pkg => github.com/rancher/pkg v0.0.0-20190514055449-b30ab9de040e
	github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
	k8s.io/client-go => k8s.io/client-go v0.18.0
)

require (
	github.com/olekukonko/tablewriter v0.0.2
	github.com/rancher/types v0.0.0-20200625174156-fe03f32597d2
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	k8s.io/api v0.18.0
	k8s.io/apimachinery v0.18.0
	k8s.io/client-go v12.0.0+incompatible
)
