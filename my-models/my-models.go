package mymodel

import "C"

import (
	//	"fmt"
	"github.com/sjwhitworth/golearn/base"
	//	"unsafe"
)

// structure

type MyClassifier struct {
	base.BaseClassifier
	bias float64
}

// New
func NewMyClassifier(prune float64) (*MyClassifier, error) {
	return &MyClassifier{
		base.BaseClassifier{},
		0.01,
	}, nil
}

// Fit
func (t *MyClassifier) Fit(on base.FixedDataGrid) error {
	return nil
}

// Predict issues predictions from a trained LinearSVC.
func (lr *MyClassifier) Predict(X base.FixedDataGrid) (base.FixedDataGrid, error) {
	return nil, nil
}

// String return a humaan-readable version.
func (lr *MyClassifier) String() string {
	return "MingClassifier"
}
