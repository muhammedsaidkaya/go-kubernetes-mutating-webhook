package main

import (
	"github.com/muhammedkaya/tcpdump-webhook/pkg/httpClient"
	"github.com/muhammedkaya/tcpdump-webhook/pkg/k8sConfig"
)

func main() {
	k8sConfig.SetK8sConfig()
	httpClient.CreateWebHookServer()
}
