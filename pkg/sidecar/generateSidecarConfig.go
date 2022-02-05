package sidecar

import (
	"github.com/muhammedkaya/tcpdump-webhook/pkg"
	corev1 "k8s.io/api/core/v1"
)

func GenerateSidecarConfig(authConfig pkg.AuthenticatorSidecarConfig) *pkg.PatchConfig {
	var containers, initContainers []corev1.Container

	sidecarContainer := corev1.Container{
		Name:            "tcpdump-sidecar",
		Image:           "bilalunalnet/tcpdump-alpine",
		ImagePullPolicy: "Always",
		Env: []corev1.EnvVar{
			{
				Name:  "CONJUR_AUTHN_TOKEN_FILE",
				Value: "/run/conjur/conjur-access-token",
			},
		},
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "conjur-access-token",
				MountPath: "/run/conjur",
			},
		},
	}

	candidates := []corev1.Container{sidecarContainer}

	if authConfig.ContainerMode == "init" {
		initContainers = candidates
	} else {
		containers = candidates
	}

	return &pkg.PatchConfig{
		Containers:     containers,
		InitContainers: initContainers,
		Volumes: []corev1.Volume{
			{
				Name: "conjur-access-token",
				VolumeSource: corev1.VolumeSource{
					EmptyDir: &corev1.EmptyDirVolumeSource{
						Medium: "Memory",
					},
				},
			},
		},
	}
}
