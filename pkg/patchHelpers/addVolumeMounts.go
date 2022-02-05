package patchHelpers

import (
	"fmt"
	"github.com/muhammedkaya/tcpdump-webhook/pkg"
	corev1 "k8s.io/api/core/v1"
)

func AddVolumeMounts(
	target []corev1.Container,
	added pkg.ContainerVolumeMounts,
	basePath string,
) (patch []pkg.PatchOperation) {
	for index, container := range target {
		volumeMounts, ok := added[container.Name]
		if !ok || len(volumeMounts) == 0 {
			continue
		}

		if len(container.VolumeMounts) == 0 {
			volumeMount := volumeMounts[0]
			volumeMounts = volumeMounts[1:]

			path := fmt.Sprintf("%s/%d/volumeMounts", basePath, index)
			patch = append(patch, pkg.PatchOperation{
				Op:    pkg.PatchOperationAdd,
				Path:  path,
				Value: []corev1.VolumeMount{volumeMount},
			})
		}

		path := fmt.Sprintf("%s/%d/volumeMounts/-", basePath, index)
		for _, volumeMount := range volumeMounts {
			patch = append(patch, pkg.PatchOperation{
				Op:    pkg.PatchOperationAdd,
				Path:  path,
				Value: volumeMount,
			})
		}
	}

	return patch
}
