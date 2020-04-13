package test

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
	core "k8s.io/api/core/v1"
	v1 "k8s.io/api/storage/v1"
	"testing"
)

var tplStorageclass = []string{"templates/storageclass.yaml"}

func Test_Storageclass_GivenClassesEnabled_WhenNoPolicyDefined_ThenRenderDefault(t *testing.T) {
	options := &helm.Options{
		SetValues: map[string]string{
			"storageClass.create": "true",
			"storageClass.classes[0].policy": "",
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplStorageclass)

	var class v1.StorageClass
	helm.UnmarshalK8SYaml(t, output, &class)

	expectedPolicy := core.PersistentVolumeReclaimDelete
	assert.Equal(t, &expectedPolicy, class.ReclaimPolicy)
}

func Test_Secret_GivenClassesEnabled_WhenNoTypeDefined_ThenRenderDefault(t *testing.T) {
	options := &helm.Options{
		SetValues: map[string]string{
			"storageClass.create": "true",
			"storageClass.classes[0].type": "",
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplStorageclass)

	var class v1.StorageClass
	helm.UnmarshalK8SYaml(t, output, &class)

	assert.Equal(t, "nfs", class.Parameters["type"])
}

func Test_Secret_GivenClassesEnabled_WhenAdditionalParametersUndefined_ThenRenderEmptyValues(t *testing.T) {
	options := &helm.Options{
		SetValues: map[string]string{
			"storageClass.create": "true",
			"storageClass.classes[0].node": "",
			"storageClass.classes[0].shareProperties": "",
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplStorageclass)

	var class v1.StorageClass
	helm.UnmarshalK8SYaml(t, output, &class)

	assert.Equal(t, "", class.Parameters["shareProperties"])
	assert.Equal(t, "", class.Parameters["node"])
}
