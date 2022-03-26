package testdata

import _ "embed"

var (
	//go:embed action.yml
	Action []byte

	//go:embed checkout_step.yml
	CheckoutStep []byte

	//go:embed checkout_test_build_workflow.yml
	CheckoutTestBuildWorkflow []byte

	//go:embed checkout_test_job.yml
	CheckoutTestJob []byte

	//go:embed default_image_step.yml
	DefaultImageStep []byte

	//go:embed docker_step.yml
	DockerStep []byte

	//go:embed env_job.yml
	EnvJob []byte

	//go:embed env_step.yml
	EnvStep []byte

	//go:embed expand_step.yml
	ExpandStep []byte

	//go:embed job_container_image_job.yml
	JobContainerImageJob []byte

	//go:embed job_container_job.yml
	JobContainerJob []byte

	//go:embed set_output_job.yml
	SetOutputJob []byte

	//go:embed stop_commands_step.yml
	StopCommandsStep []byte
)
