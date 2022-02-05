package pkg

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	UniversalDeserializer = serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer()
)
var Config *rest.Config
var ClientSet *kubernetes.Clientset
var Parameters ServerParameters

const (
	PatchOperationAdd     = "add"
	PatchOperationReplace = "replace"
)
