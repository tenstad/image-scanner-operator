// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/statnett/image-scanner-operator/api/stas/v1alpha1"
	stasv1alpha1 "github.com/statnett/image-scanner-operator/internal/client/applyconfiguration/stas/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=stas.statnett.no, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("ContainerImageScan"):
		return &stasv1alpha1.ContainerImageScanApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ContainerImageScanSpec"):
		return &stasv1alpha1.ContainerImageScanSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ContainerImageScanStatus"):
		return &stasv1alpha1.ContainerImageScanStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Image"):
		return &stasv1alpha1.ImageApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageScanSpec"):
		return &stasv1alpha1.ImageScanSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ScanConfig"):
		return &stasv1alpha1.ScanConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Vulnerability"):
		return &stasv1alpha1.VulnerabilityApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("VulnerabilitySummary"):
		return &stasv1alpha1.VulnerabilitySummaryApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Workload"):
		return &stasv1alpha1.WorkloadApplyConfiguration{}

	}
	return nil
}