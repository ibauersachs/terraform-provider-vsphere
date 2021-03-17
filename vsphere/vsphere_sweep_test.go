package vsphere

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestMain(m *testing.M) {
	acctest.UseBinaryDriver("vsphere", Provider)
	resource.TestMain(m)
}

func sweepVSphereClient() (*VSphereClient, error) {
	config := Config{
		InsecureFlag:    true,
		Debug:           false,
		Persist:         true,
		User:            os.Getenv("TF_VAR_VSPHERE_USER"),
		Password:        os.Getenv("TF_VAR_VSPHERE_PASSWORD"),
		VSphereServer:   os.Getenv("TF_VAR_VSPHERE_SERVER"),
		DebugPath:       "",
		DebugPathRun:    "",
		VimSessionPath:  "",
		RestSessionPath: "",
		KeepAlive:       0,
	}
	return config.Client()
}
