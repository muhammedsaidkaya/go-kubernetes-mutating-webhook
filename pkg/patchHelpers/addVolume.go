package patchHelpers

import (
	"github.com/muhammedkaya/tcpdump-webhook/pkg"
	corev1 "k8s.io/api/core/v1"
)

func AddVolume(
	target, added []corev1.Volume,
	basePath string,
) (patch []pkg.PatchOperation) {
	first := len(target) == 0
	var value interface{}

	for _, add := range added {
		value = add
		path := basePath

		if first {
			first = false
			value = []corev1.Volume{add}
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
