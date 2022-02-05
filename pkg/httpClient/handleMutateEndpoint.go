package httpClient

import (
	"encoding/json"
	"fmt"
	"github.com/muhammedkaya/tcpdump-webhook/pkg/sidecar"
	apiv1 "k8s.io/api/core/v1"
	"net/http"
)

func HandleMutateEndpoint(w http.ResponseWriter, r *http.Request) {

	admissionReviewReq := CastHttpBodyToAdmssionReviewRequest(w, r)
	var pod apiv1.Pod

	err := json.Unmarshal(admissionReviewReq.Request.Object.Raw, &pod)
	if err != nil {
		fmt.Errorf("could not unmarshal pod on admission request: %v", err)
	}

	patchBytes, err := sidecar.CreateSidecarPatch(pod)
	if err != nil {
		fmt.Errorf("could not marshal JSON patch: %v", err)
	}
	CreateAdmissionReviewResponse(admissionReviewReq, patchBytes, w)

}
