package controller

import (
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func getAnnotationValue(resource *unstructured.Unstructured, annotationKey string) (map[string]interface{}, error) {
	//Get the annotations from the resource
	annotations := resource.GetAnnotations()

	// Check if the annotation does not exist
	jsonStr, exists := annotations[annotationKey]
	if !exists {
		fmt.Printf("Annotation %s does not exist\n", annotationKey)
		return nil, nil
	}

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return nil, err
	}

	fmt.Printf("Parsed JSON: %v\n", data)
	return data, nil
}

func setJSONAnnotation(resource *unstructured.Unstructured, annotationKey string, anotationData map[string]interface{}) error {
	// Create the JSON data
	data := createAnnotationData()

	// Serialize to JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Add annotation
	annotations := resource.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[annotationKey] = string(jsonData)
	resource.SetAnnotations(annotations)

	return nil
}

func createAnnotationData() map[string]interface{} {
	data := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]string{
			"nestedKey": "nestedValue",
		},
		"key3": []int{1, 2, 3},
	}
	return data
}
