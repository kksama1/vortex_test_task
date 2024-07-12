// Package pod_placeholder provides "Pods" placeholder needed to solve the task.
package pod_placeholder

import (
	"errors"
	"log"
)

type Deployer interface {
	CreatePod(name string) error
	DeletePod(name string) error
	GetPodList() ([]string, error)
}

type PodList struct {
	Pods []PodPlaceholder
}

type PodPlaceholder struct {
	name string
}

// CreatePod method creates new pod and appends it to pod list.
func (p *PodList) CreatePod(name string) error {
	for i := range p.Pods {
		if p.Pods[i].name == name {
			return errors.New("pod already exists")
		}
	}
	pod := PodPlaceholder{name: name}
	p.Pods = append(p.Pods, pod)
	log.Printf("Pod created: %s\n", name)
	return nil
}

// DeletePod method deletes pod and reduces pod list.
func (p *PodList) DeletePod(name string) error {
	for i := range p.Pods {
		if p.Pods[i].name == name {
			p.Pods = append(p.Pods[:i], p.Pods[i+1:]...)
			log.Printf("Deleted pod: %s\n", name)
			return nil
		}
	}
	return nil
}

// GetPodList returns PodPlaceholder slice as string slice representation.
func (p *PodList) GetPodList() ([]string, error) {
	var podsNames []string
	for i := range p.Pods {
		podsNames = append(podsNames, p.Pods[i].name)
	}
	return podsNames, nil
}
