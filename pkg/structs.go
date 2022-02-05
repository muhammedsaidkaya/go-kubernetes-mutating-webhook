package pkg

import (
	corev1 "k8s.io/api/core/v1"
)

type ContainerVolumeMounts map[string][]corev1.VolumeMount

type PatchConfig struct {
	InitContainers        []corev1.Container    `yaml:"initContainers"`
	Containers            []corev1.Container    `yaml:"containers"`
	Volumes               []corev1.Volume       `yaml:"volumes"`
	ContainerVolumeMounts ContainerVolumeMounts `yaml:"volumeMounts"`
}

type AuthenticatorSidecarConfig struct {
	ContainerMode string
}

type ServerParameters struct {
	Port     int    // webhook server port
	CertFile string // path to the x509 certificate for https
	KeyFile  string // path to the x509 private key matching `CertFile`
}

type PatchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}
