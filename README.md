# Eraser: Scanner Template

This repo is to be used with [Eraser](https://github.com/Azure/eraser), as a template for the customizable scanner.

To use, implement the `scan()` function, which takes in a list of all non-running, non-excluded images in the cluster. This function should return any non-compliant images and any images that failed to scan. When complete, provide your custom scanner image to Eraser in deployment.
