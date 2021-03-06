package envinfo

import (
	"fmt"

	devfilev1 "github.com/devfile/api/pkg/apis/workspaces/v1alpha2"
	"github.com/openshift/odo/pkg/localConfigProvider"
)

const (
	// DefaultVolumeSize Default volume size for volumes defined in a devfile
	DefaultVolumeSize = "1Gi"
)

// CompleteStorage completes the given storage
func (ei *EnvInfo) CompleteStorage(storage *localConfigProvider.LocalStorage) {
	if storage.Size == "" {
		storage.Size = DefaultVolumeSize
	}
	if storage.Path == "" {
		// acc to the devfile schema, if the mount path is absent; it will be mounted at the dir with the mount name
		storage.Path = "/" + storage.Name
	}
}

// ValidateStorage validates the given storage
func (ei *EnvInfo) ValidateStorage(storage localConfigProvider.LocalStorage) error {
	for _, store := range ei.ListStorage() {
		if store.Name == storage.Name {
			return fmt.Errorf("storage with name %s already exists", storage.Name)
		}
	}
	return nil
}

// GetStorage gets the storage with the given name
func (ei *EnvInfo) GetStorage(name string) *localConfigProvider.LocalStorage {
	for _, storage := range ei.ListStorage() {
		if name == storage.Name {
			return &storage
		}
	}
	return nil
}

// CreateStorage sets the storage related information in the local configuration
func (ei *EnvInfo) CreateStorage(storage localConfigProvider.LocalStorage) error {
	err := ei.devfileObj.Data.AddVolume(devfilev1.Component{
		Name: storage.Name,
		ComponentUnion: devfilev1.ComponentUnion{
			Volume: &devfilev1.VolumeComponent{
				Volume: devfilev1.Volume{
					Size: storage.Size,
				},
			},
		},
	}, storage.Path)

	if err != nil {
		return err
	}
	err = ei.devfileObj.WriteYamlDevfile()
	if err != nil {
		return err
	}

	return nil
}

// ListStorage gets all the storage from the devfile.yaml
func (ei *EnvInfo) ListStorage() []localConfigProvider.LocalStorage {
	volumeSizeMap := make(map[string]string)
	components := ei.devfileObj.Data.GetComponents()

	for _, component := range components {
		if component.Volume == nil {
			continue
		}
		if component.Volume.Size == "" {
			component.Volume.Size = DefaultVolumeSize
		}
		volumeSizeMap[component.Name] = component.Volume.Size
	}

	var storageList []localConfigProvider.LocalStorage
	for _, component := range components {
		if component.Container == nil {
			continue
		}
		for _, volumeMount := range component.Container.VolumeMounts {
			size, ok := volumeSizeMap[volumeMount.Name]
			if ok {
				storageList = append(storageList, localConfigProvider.LocalStorage{
					Name:      volumeMount.Name,
					Size:      size,
					Path:      GetVolumeMountPath(volumeMount),
					Container: component.Name,
				})
			}
		}
	}

	return storageList
}

// DeleteStorage deletes the storage with the given name
func (ei *EnvInfo) DeleteStorage(name string) error {
	err := ei.devfileObj.Data.DeleteVolume(name)
	if err != nil {
		return err
	}
	err = ei.devfileObj.WriteYamlDevfile()
	if err != nil {
		return err
	}

	return nil
}

// GetStorageMountPath gets the mount path of the storage with the given storage name
func (ei *EnvInfo) GetStorageMountPath(storageName string) (string, error) {
	return ei.devfileObj.Data.GetVolumeMountPath(storageName)
}

// GetVolumeMountPath gets the volume mount's path
func GetVolumeMountPath(volumeMount devfilev1.VolumeMount) string {
	// if there is no volume mount path, default to volume mount name as per devfile schema
	if volumeMount.Path == "" {
		volumeMount.Path = "/" + volumeMount.Name
	}

	return volumeMount.Path
}
