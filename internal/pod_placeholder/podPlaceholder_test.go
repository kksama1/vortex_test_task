package pod_placeholder

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCreatePod(t *testing.T) {
	pod2 := PodPlaceholder{name: "Pod № 2"}
	pod3 := PodPlaceholder{"Pod № 3"}
	podlist := PodList{Pods: []PodPlaceholder{pod2, pod3}}
	tests := []struct {
		name     string
		podName  string
		expected error
	}{
		{
			name:     "create pod",
			podName:  "Pod № 1",
			expected: nil,
		},
		{
			name:     "pod already exists",
			podName:  "Pod № 2",
			expected: errors.New("pod already exists"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected == nil {
				assert.NoError(t, podlist.CreatePod(tt.podName), tt.expected)
			} else {
				assert.Error(t, podlist.CreatePod(tt.podName), tt.expected)
			}
		})
	}
}

func TestDeletePod(t *testing.T) {
	pod2 := PodPlaceholder{name: "Pod № 2"}
	pod3 := PodPlaceholder{"Pod № 3"}
	podlist := PodList{Pods: []PodPlaceholder{pod2, pod3}}
	tests := []struct {
		name     string
		podName  string
		expected error
	}{
		{
			name:     "no such pod",
			podName:  "Pod № 1",
			expected: nil,
		},
		{
			name:     "delete existing pod",
			podName:  "Pod № 2",
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected == nil {
				assert.NoError(t, podlist.DeletePod(tt.podName), tt.expected)
			} else {
				assert.Error(t, podlist.CreatePod(tt.podName), tt.expected)
			}
		})
	}
}

func TestPodList(t *testing.T) {
	pod2 := PodPlaceholder{name: "Pod № 2"}
	pod3 := PodPlaceholder{"Pod № 3"}
	podlist := PodList{Pods: []PodPlaceholder{pod2, pod3}}
	expectedRes := []string{"Pod № 2", "Pod № 3"}

	t.Run("Struct to string (?)", func(t *testing.T) {
		res, err := podlist.GetPodList()
		reflect.DeepEqual(res, expectedRes)
		assert.NoError(t, err)
	})

}
