/*
Copyright(c) 2023-present Accton. All rights reserved. www.accton.com
*/

package virtualization

import (
	"context"
	"strconv"
	"testing"

	fakek8s "k8s.io/client-go/kubernetes/fake"
	fakeks "kubesphere.io/kubesphere/pkg/client/clientset/versioned/fake"

	vm_ctrl "kubesphere.io/kubesphere/pkg/controller/virtualization/virtualmachine"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	cdiv1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	virtzv1alpha1 "kubesphere.io/api/virtualization/v1alpha1"
	ui_virtz "kubesphere.io/kubesphere/pkg/models/virtualization"
)

func prepareFakeDiskVolume(ksClient *fakeks.Clientset, vm_instance *virtzv1alpha1.VirtualMachine) error {

	for _, diskVolumeTemplate := range vm_instance.Spec.DiskVolumeTemplates {
		disk := vm_ctrl.GenerateDiskVolume(vm_instance, &diskVolumeTemplate)

		_, err := ksClient.VirtualizationV1alpha1().DiskVolumes(vm_instance.Namespace).Create(context.Background(), disk, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

type FakeImageTemplate struct {
	Name      string
	Namespace string
	Size      uint
}

func prepareFakeImageTemplate(ksClient *fakeks.Clientset, fakeImageTemlate *FakeImageTemplate) error {
	imageName := fakeImageTemlate.Name
	imageNamespace := fakeImageTemlate.Namespace

	size := strconv.FormatUint(uint64(fakeImageTemlate.Size), 10)
	imagetemplate := &virtzv1alpha1.ImageTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      imageName,
			Namespace: imageNamespace,
		},
		Spec: virtzv1alpha1.ImageTemplateSpec{
			Source: virtzv1alpha1.ImageTemplateSource{
				HTTP: &cdiv1.DataVolumeSourceHTTP{
					URL: "http://test.com",
				},
			},
			Resources: virtzv1alpha1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceStorage: resource.MustParse(size + "Gi"),
				},
			},
		},
	}

	_, err := ksClient.VirtualizationV1alpha1().ImageTemplates(fakeImageTemlate.Namespace).Create(context.TODO(), imagetemplate, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func createVirtualMachine(h *virtzhandler, ui_virtz_vm *ui_virtz.VirtualMachineRequest, namespace string) (*virtzv1alpha1.VirtualMachine, error) {
	vm, err := h.virtz.CreateVirtualMachine(namespace, ui_virtz_vm)
	if err != nil {
		return nil, err
	}

	return vm, nil
}

func createImage(h *virtzhandler, ui_virtz_image *ui_virtz.ImageRequest, namespace string) (*virtzv1alpha1.ImageTemplate, error) {
	image, err := h.virtz.CreateImage(namespace, ui_virtz_image)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func createDisk(h *virtzhandler, ui_virtz_disk *ui_virtz.DiskRequest, namespace string) (*virtzv1alpha1.DiskVolume, error) {
	disk, err := h.virtz.CreateDisk(namespace, ui_virtz_disk)
	if err != nil {
		return nil, err
	}

	return disk, nil
}

func prepareFakeMinioService(k8sClient *fakek8s.Clientset) error {
	minioService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "minio",
			Namespace: "default",
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Name: "minio",
					Port: 9000,
				},
			},
			ClusterIP: "1.2.3.4",
		},
	}

	_, err := k8sClient.CoreV1().Services("default").Create(context.Background(), minioService, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func checkVirtualMachineResult(t *testing.T, vm *virtzv1alpha1.VirtualMachine, ui_vm_req ui_virtz.VirtualMachineRequest) {
	if vm.Annotations[virtzv1alpha1.VirtualizationAliasName] != ui_vm_req.Name {
		t.Errorf("vm alias name is not correct: got %v want %v", vm.Annotations[virtzv1alpha1.VirtualizationAliasName], ui_vm_req.Name)
	}

	if vm.Spec.Hardware.Domain.CPU.Cores != uint32(ui_vm_req.CpuCores) {
		t.Errorf("vm cpu cores is not correct: got %v want %v", vm.Spec.Hardware.Domain.CPU.Cores, ui_vm_req.CpuCores)
	}

	memory := strconv.FormatUint(uint64(ui_vm_req.Memory), 10) + "Gi"
	if vm.Spec.Hardware.Domain.Resources.Requests.Memory().String() != memory {
		t.Errorf("vm memory is not correct: got %v want %v", vm.Spec.Hardware.Domain.Resources.Requests.Memory().String(), memory)
	}

	// check only system disk
	if len(vm.Spec.DiskVolumes) == 1 {
		if len(vm.Spec.DiskVolumeTemplates) != 1 {
			t.Errorf("vm disk number is not correct: got %v want %v", len(vm.Spec.DiskVolumeTemplates), 1)
		}

		if vm.Spec.DiskVolumeTemplates[0].Labels[virtzv1alpha1.VirtualizationDiskType] != "system" {
			t.Errorf("vm disk type is not correct: got %v want %v", vm.Spec.DiskVolumeTemplates[0].Labels[virtzv1alpha1.VirtualizationDiskType], "system")
		}

		system_size := strconv.FormatUint(uint64(ui_vm_req.Image.Size), 10) + "Gi"
		if vm.Spec.DiskVolumeTemplates[0].Spec.Resources.Requests.Storage().String() != system_size {
			t.Errorf("vm disk size is not correct: got %v want %v", vm.Spec.DiskVolumeTemplates[0].Spec.Resources.Requests.Storage().String(), system_size)
		}
	}

	// check system and data disk
	if len(vm.Spec.DiskVolumes) == 2 {
		if len(ui_vm_req.Disk) == 1 {
			if ui_vm_req.Disk[0].Action == "add" || ui_vm_req.Disk[0].Action == "mount" {
				if len(vm.Spec.DiskVolumeTemplates) != 2 {
					t.Errorf("vm disk number is not correct: got %v want %v", len(vm.Spec.DiskVolumeTemplates), 2)
				}

				if vm.Spec.DiskVolumeTemplates[1].Labels[virtzv1alpha1.VirtualizationDiskType] != "data" {
					t.Errorf("vm disk type is not correct: got %v want %v", vm.Spec.DiskVolumeTemplates[1].Labels[virtzv1alpha1.VirtualizationDiskType], "data")
				}

				data_size := strconv.FormatUint(uint64(ui_vm_req.Disk[0].Size), 10) + "Gi"
				if vm.Spec.DiskVolumeTemplates[1].Spec.Resources.Requests.Storage().String() != data_size {
					t.Errorf("vm disk size is not correct: got %v want %v", vm.Spec.DiskVolumeTemplates[1].Spec.Resources.Requests.Storage().String(), data_size)
				}
			}
		}
	}

	for _, uiDisk := range ui_vm_req.Disk {
		if uiDisk.Action == "add" || uiDisk.Action == "mount" {
			if len(vm.Spec.DiskVolumes) != len(ui_vm_req.Disk)+1 {
				t.Errorf("vm disk number is not correct: got %v want %v", len(vm.Spec.DiskVolumes), len(ui_vm_req.Disk)+1)
			}
		}
	}
}
