package controller

import (
	"context"
	"net/http"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	jsonpatch "gomodules.xyz/jsonpatch/v2"
)

type PodAnnotator struct {
	Client  client.Client
	Logger  logr.Logger
	decoder *admission.Decoder
}

// podAnnotator adds an annotation to every incoming pods.
func (a *PodAnnotator) Handle(ctx context.Context, req admission.Request) admission.Response {
	object := &metav1.PartialObjectMetadata{}
	err := a.decoder.Decode(req, object)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	patches := []jsonpatch.JsonPatchOperation{}
	if object.ObjectMeta.Labels == nil {
		patches = append(patches, jsonpatch.NewOperation("add", "/metadata/labels", struct{}{}))
		object.ObjectMeta.Labels = make(map[string]string)
	}

	if _, ok := object.ObjectMeta.Labels["foo"]; !ok {
		patches = append(patches, jsonpatch.NewOperation("add", "/metadata/labels/foo", "webhook-was-here"))
	}

	var message string
	if len(patches) > 0 {
		message = "Added Label \"foo\", because it was missing."
	} else {
		message = "Doing nothing. Label \"foo\" was already present."
	}

	return admission.Patched(message, patches...)
}

// podAnnotator implements admission.DecoderInjector.
// A decoder will be automatically injected.

// InjectDecoder injects the decoder.
func (a *PodAnnotator) InjectDecoder(d *admission.Decoder) error {
	a.decoder = d
	return nil
}

func (a *PodAnnotator) InjectLogger(l logr.Logger) error {
	a.Logger = l
	return nil
}
