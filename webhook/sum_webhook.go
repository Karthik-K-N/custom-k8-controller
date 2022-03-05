package webhook

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	v1 "github.com/Karthik-K-N/custom-k8-controller/api/v1"
)

type SumWebhook struct {
}

func (r *SumWebhook) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(&v1.Sum{}).
		WithValidator(r).
		Complete()
}

//+kubebuilder:webhook:path=/validate-calculator-sample-domain-v1-sum,mutating=false,failurePolicy=fail,sideEffects=None,groups=calculator.sample.domain,resources=sums,verbs=create;update,versions=v1,name=vsum.kb.io,admissionReviewVersions=v1

var _ webhook.CustomValidator = &SumWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *SumWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	sum, ok := obj.(*v1.Sum)
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("expected a Sum but got a %T", obj))
	}
	if sum.Spec.NumberOne < 0 || sum.Spec.NumberTwo < 0 {
		return fmt.Errorf("The input values Number One: %d Number Two: %d cannot be negative ", sum.Spec.NumberOne, sum.Spec.NumberTwo)
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *SumWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *SumWebhook) ValidateDelete(_ context.Context, obj runtime.Object) error {
	return nil
}
