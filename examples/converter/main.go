package main

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api/legacyscheme"

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/apis/apps"

	_ "k8s.io/kubernetes/pkg/apis/apps/install"
)

func main() {
	v1beta1Deployment := &appsv1beta1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1beta1",
		},
	}
	fmt.Println("obj v1beta1: ", v1beta1Deployment.GetObjectKind().GroupVersionKind().String())

	objInternal, err := legacyscheme.Scheme.ConvertToVersion(v1beta1Deployment, apps.SchemeGroupVersion)
	utilruntime.Must(err)

	fmt.Println("obj internal: ", objInternal.GetObjectKind().GroupVersionKind().String())

	objV1beta2, err := legacyscheme.Scheme.ConvertToVersion(objInternal, appsv1beta2.SchemeGroupVersion)
	utilruntime.Must(err)

	fmt.Println("obj v1beta2: ", objV1beta2.GetObjectKind().GroupVersionKind().String())
}
