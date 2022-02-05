package sidecar

import (
	"encoding/json"
	"github.com/muhammedkaya/tcpdump-webhook/pkg"
	"github.com/muhammedkaya/tcpdump-webhook/pkg/patchHelpers"
	corev1 "k8s.io/api/core/v1"
)

func createPatch(
	pod *corev1.Pod,
	sidecarConfig *pkg.PatchConfig,
	annotations map[string]string,
) ([]byte, error) {
	var patch []pkg.PatchOperation

	//patch = append(
	//	patch,
	//	patchHelpers.AddContainer(
	//		pod.Spec.InitContainers,
	//		sidecarConfig.InitContainers,
	//		"/spec/initContainers",
	//	)...,
	//)
	patch = append(
		patch,
		patchHelpers.AddContainer(
			pod.Spec.Containers,
			sidecarConfig.Containers,
			"/spec/containers",
		)...,
	)
	patch = append(
		patch,
		patchHelpers.AddVolume(
			pod.Spec.Volumes,
			sidecarConfig.Volumes, "/spec/volumes",
		)...,
	)
	patch = append(
		patch,
		patchHelpers.UpdateAnnotation(
			pod.Annotations,
			annotations,
		)...,
	)
	return json.Marshal(patch)
}
