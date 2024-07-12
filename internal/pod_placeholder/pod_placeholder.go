package pod_placeholder

type Deployer interface {
	CreatePod(name string) error
	DeletePod(name string) error
	GetPodList() ([]string, error)
}

type pod_placeholder struct {
	id                int64
	algorithmHolderId int64
	userHolderId      int64
	algorithmActive   []string
	isActive          bool
}

func (p *pod_placeholder) CreatePod(name string) error {

	return nil
}
