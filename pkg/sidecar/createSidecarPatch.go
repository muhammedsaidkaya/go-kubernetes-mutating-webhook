package sidecar

import (
	"github.com/muhammedkaya/tcpdump-webhook/pkg"
	apiv1 "k8s.io/api/core/v1"
)

func CreateSidecarPatch(pod apiv1.Pod) ([]byte, error) {
	sidecarConfig := GenerateSidecarConfig(pkg.AuthenticatorSidecarConfig{
		ContainerMode: "not-init",
	})
	annotations := map[string]string{"annotationStatusKey": "injected"}
	return createPatch(&pod, sidecarConfig, annotations)
}
