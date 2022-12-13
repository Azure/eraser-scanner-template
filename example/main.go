package main

import (
	"context"

	eraserv1alpha1 "github.com/Azure/eraser/api/v1alpha1"
	template "github.com/Azure/eraser/api/v1alpha1/pkg/scanners/template"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

func main() {
	log := logf.Log.WithName("scanner").WithValues("provider", "customScanner")

	// create image provider with custom values
	imageProvider := template.NewImageProvider(
		template.WithContext(context.Background()),
		template.WithMetrics(true),
		template.WithDeleteScanFailedImages(true),
		template.WithLogger(log),
	)

	// retrieve list of all non-running, non-excluded images from collector container
	allImages, err := imageProvider.RecieveImages()
	if err != nil {
		log.Error(err, "unable to retrieve list of images from collector container")
	}

	// scan images with custom scanner
	nonCompliant, failedImages := scan(allImages)

	// send images to eraser container
	if err := imageProvider.SendImages(nonCompliant, failedImages); err != nil {
		log.Error(err, "unable to send non-compliant images to eraser container")
		return err
	}

	// complete scan
	imageProvider.Finish()
}

// TODO: implement customized scanner
func scan(allImages []eraserv1alpha1.Image) ([]eraserv1alpha1.Image, []eraserv1alpha1.Image) {
	// scan images and partition into non-compliant and failedImages
	var nonCompliant, failedImages []eraserv1alpha1.Image

	return nonCompliant, failedImages
}
