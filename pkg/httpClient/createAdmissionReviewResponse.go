package httpClient

import (
	"encoding/json"
	"fmt"
	"k8s.io/api/admission/v1beta1"
	"net/http"
)

func CreateAdmissionReviewResponse(admissionReviewReq v1beta1.AdmissionReview, patchBytes []byte, w http.ResponseWriter) {
	admissionReviewResponse := v1beta1.AdmissionReview{
		Response: &v1beta1.AdmissionResponse{
			UID:     admissionReviewReq.Request.UID,
			Allowed: true,
		},
	}

	admissionReviewResponse.Response.Patch = patchBytes

	bytes, err := json.Marshal(&admissionReviewResponse)
	if err != nil {
		fmt.Errorf("marshaling response: %v", err)
	}

	w.Write(bytes)
}
