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

	// TODO: implement customized scanner to scan allImages and  partition into vulnerableImages and failedImages
	vulnerableImages := make([]eraserv1alpha1.Image, 0, len(allImages))
	failedImages := make([]eraserv1alpha1.Image, 0, len(allImages))

	// send images to eraser container
	if err := imageProvider.SendImages(vulnerableImages, failedImages); err != nil {
		log.Error(err, "unable to send non-compliant images to eraser container")
		return err
	}

	// complete scan
	imageProvider.Finish()
}
