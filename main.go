package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	//"github.com/sjwhitworth/golearn/linear_models"
	"github.com/mingderwang/golearn/mnist/handwrittenDigitRecognition/my-models"
)

func main() {
	// Load and parse the data from csv files
	fmt.Println("Loading data...")
	/*
		trainData, err := base.ParseCSVToInstances("data/mnist_train.csv", true)
		if err != nil {
			panic(err)
		}
	*/
	testData, err := base.ParseCSVToInstances("data/mnist_test.csv", true)
	if err != nil {
		panic(err)
	}
	// Create a new linear SVC with some good default values
	classifier, err := mymodel.NewMyClassifier(1.0)
	if err != nil {
		panic(err)
	}

	// Don't output information on each iteration
	//base.Silent()

	// Train the linear SVC
	fmt.Println("Training...")
	//fmt.Println(testData)
	classifier.Fit(testData)

	//write my own classifier
}
