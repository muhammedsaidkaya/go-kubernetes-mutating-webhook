package patchHelpers

import (
	"github.com/muhammedkaya/tcpdump-webhook/pkg"
	corev1 "k8s.io/api/core/v1"
)

func AddContainer(
	target, added []corev1.Container,
	basePath string,
) (patch []pkg.PatchOperation) {
	first := len(target) == 0
	var value interface{}

	for _, add := range added {
		value = add
		path := basePath
		if first {
			first = false
			value = []corev1.Container{add}
		} else {
			path = path + "/-"
		}
		patch = append(patch, pkg.PatchOperation{
			Op:    pkg.PatchOperationAdd,
			Path:  path,
			Value: value,
		})
	}

	return patch
}
