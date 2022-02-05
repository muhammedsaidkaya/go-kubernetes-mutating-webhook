package httpClient

import (
	"errors"
	"fmt"
	"github.com/muhammedkaya/tcpdump-webhook/pkg"
	"io/ioutil"
	"k8s.io/api/admission/v1beta1"
	"net/http"
)

func CastHttpBodyToAdmssionReviewRequest(w http.ResponseWriter, r *http.Request) (admissionReviewReq v1beta1.AdmissionReview) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("could not read request: %v", err)
	}

	if _, _, err := pkg.UniversalDeserializer.Decode(body, nil, &admissionReviewReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("could not deserialize request: %v", err)
	} else if admissionReviewReq.Request == nil {
		w.WriteHeader(http.StatusBadRequest)
		errors.New("malformed admission review: request is nil")
	}
	return
}
