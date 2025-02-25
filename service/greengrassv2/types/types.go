// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Contains a request to associate a client device with a core device. The
// BatchAssociateClientDeviceWithCoreDevice
// (https://docs.aws.amazon.com/greengrass/v2/APIReference/API_BatchAssociateClientDeviceWithCoreDevice.html)
// operation consumes a list of these requests.
type AssociateClientDeviceWithCoreDeviceEntry struct {

	// The name of the IoT thing that represents the client device to associate.
	//
	// This member is required.
	ThingName *string

	noSmithyDocumentSerde
}

// Contains an error that occurs from a request to associate a client device with a
// core device. The BatchAssociateClientDeviceWithCoreDevice
// (https://docs.aws.amazon.com/greengrass/v2/APIReference/API_BatchAssociateClientDeviceWithCoreDevice.html)
// operation returns a list of these errors.
type AssociateClientDeviceWithCoreDeviceErrorEntry struct {

	// The error code for the request.
	Code *string

	// A message that provides additional information about the error.
	Message *string

	// The name of the IoT thing whose associate request failed.
	ThingName *string

	noSmithyDocumentSerde
}

// Contains information about a client device that is associated to a core device
// for cloud discovery.
type AssociatedClientDevice struct {

	// The time that the client device was associated, expressed in ISO 8601 format.
	AssociationTimestamp *time.Time

	// The name of the IoT thing that represents the associated client device.
	ThingName *string

	noSmithyDocumentSerde
}

// Contains the status of a component version in the IoT Greengrass service.
type CloudComponentStatus struct {

	// The state of the component version.
	ComponentState CloudComponentState

	// A dictionary of errors that communicate why the component version is in an error
	// state. For example, if IoT Greengrass can't access an artifact for the component
	// version, then errors contains the artifact's URI as a key, and the error message
	// as the value for that key.
	Errors map[string]string

	// A message that communicates details, such as errors, about the status of the
	// component version.
	Message *string

	// The vendor guidance state for the component version. This state indicates
	// whether the component version has any issues that you should consider before you
	// deploy it. The vendor guidance state can be:
	//
	// * ACTIVE – This component version
	// is available and recommended for use.
	//
	// * DISCONTINUED – This component version
	// has been discontinued by its publisher. You can deploy this component version,
	// but we recommend that you use a different version of this component.
	//
	// * DELETED
	// – This component version has been deleted by its publisher, so you can't deploy
	// it. If you have any existing deployments that specify this component version,
	// those deployments will fail.
	VendorGuidance VendorGuidance

	// A message that communicates details about the vendor guidance state of the
	// component version. This message communicates why a component version is
	// discontinued or deleted.
	VendorGuidanceMessage *string

	noSmithyDocumentSerde
}

// Contains information about a component.
type Component struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the component version.
	Arn *string

	// The name of the component.
	ComponentName *string

	// The latest version of the component and its details.
	LatestVersion *ComponentLatestVersion

	noSmithyDocumentSerde
}

// Contains information about a component that is a candidate to deploy to a
// Greengrass core device.
type ComponentCandidate struct {

	// The name of the component.
	ComponentName *string

	// The version of the component.
	ComponentVersion *string

	// The version requirements for the component's dependencies. Greengrass core
	// devices get the version requirements from component recipes. IoT Greengrass V2
	// uses semantic version constraints. For more information, see Semantic Versioning
	// (https://semver.org/).
	VersionRequirements map[string]string

	noSmithyDocumentSerde
}

// Contains information about a deployment's update to a component's configuration
// on Greengrass core devices. For more information, see Update component
// configurations
// (https://docs.aws.amazon.com/greengrass/v2/developerguide/update-component-configurations.html)
// in the IoT Greengrass V2 Developer Guide.
type ComponentConfigurationUpdate struct {

	// A serialized JSON string that contains the configuration object to merge to
	// target devices. The core device merges this configuration with the component's
	// existing configuration. If this is the first time a component deploys on a
	// device, the core device merges this configuration with the component's default
	// configuration. This means that the core device keeps it's existing configuration
	// for keys and values that you don't specify in this object. For more information,
	// see Merge configuration updates
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/update-component-configurations.html#merge-configuration-update)
	// in the IoT Greengrass V2 Developer Guide.
	Merge *string

	// The list of configuration nodes to reset to default values on target devices.
	// Use JSON pointers to specify each node to reset. JSON pointers start with a
	// forward slash (/) and use forward slashes to separate the key for each level in
	// the object. For more information, see the JSON pointer specification
	// (https://tools.ietf.org/html/rfc6901) and Reset configuration updates
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/update-component-configurations.html#reset-configuration-update)
	// in the IoT Greengrass V2 Developer Guide.
	Reset []string

	noSmithyDocumentSerde
}

// Contains information about a component dependency for a Lambda function
// component.
type ComponentDependencyRequirement struct {

	// The type of this dependency. Choose from the following options:
	//
	// * SOFT – The
	// component doesn't restart if the dependency changes state.
	//
	// * HARD – The
	// component restarts if the dependency changes state.
	//
	// Default: HARD
	DependencyType ComponentDependencyType

	// The component version requirement for the component dependency. IoT Greengrass
	// V2 uses semantic version constraints. For more information, see Semantic
	// Versioning (https://semver.org/).
	VersionRequirement *string

	noSmithyDocumentSerde
}

// Contains information about a component to deploy.
type ComponentDeploymentSpecification struct {

	// The version of the component.
	ComponentVersion *string

	// The configuration updates to deploy for the component. You can define reset
	// updates and merge updates. A reset updates the keys that you specify to the
	// default configuration for the component. A merge updates the core device's
	// component configuration with the keys and values that you specify. The IoT
	// Greengrass Core software applies reset updates before it applies merge updates.
	// For more information, see Update component configurations
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/update-component-configurations.html)
	// in the IoT Greengrass V2 Developer Guide.
	ConfigurationUpdate *ComponentConfigurationUpdate

	// The system user and group that the IoT Greengrass Core software uses to run
	// component processes on the core device. If you omit this parameter, the IoT
	// Greengrass Core software uses the system user and group that you configure for
	// the core device. For more information, see Configure the user and group that run
	// components
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-greengrass-core-v2.html#configure-component-user)
	// in the IoT Greengrass V2 Developer Guide.
	RunWith *ComponentRunWith

	noSmithyDocumentSerde
}

// Contains information about the latest version of a component.
type ComponentLatestVersion struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the component version.
	Arn *string

	// The version of the component.
	ComponentVersion *string

	// The time at which the component was created, expressed in ISO 8601 format.
	CreationTimestamp *time.Time

	// The description of the component version.
	Description *string

	// The platforms that the component version supports.
	Platforms []ComponentPlatform

	// The publisher of the component version.
	Publisher *string

	noSmithyDocumentSerde
}

// Contains information about a platform that a component supports.
type ComponentPlatform struct {

	// A dictionary of attributes for the platform. The IoT Greengrass Core software
	// defines the os and platform by default. You can specify additional platform
	// attributes for a core device when you deploy the Greengrass nucleus component.
	// For more information, see the Greengrass nucleus component
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html)
	// in the IoT Greengrass V2 Developer Guide.
	Attributes map[string]string

	// The friendly name of the platform. This name helps you identify the platform. If
	// you omit this parameter, IoT Greengrass creates a friendly name from the os and
	// architecture of the platform.
	Name *string

	noSmithyDocumentSerde
}

// Contains information system user and group that the IoT Greengrass Core software
// uses to run component processes on the core device. For more information, see
// Configure the user and group that run components
// (https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-greengrass-core-v2.html#configure-component-user)
// in the IoT Greengrass V2 Developer Guide.
type ComponentRunWith struct {

	// The POSIX system user and, optionally, group to use to run this component on
	// Linux core devices. The user, and group if specified, must exist on each Linux
	// core device. Specify the user and group separated by a colon (:) in the
	// following format: user:group. The group is optional. If you don't specify a
	// group, the IoT Greengrass Core software uses the primary user for the group. If
	// you omit this parameter, the IoT Greengrass Core software uses the default
	// system user and group that you configure on the Greengrass nucleus component.
	// For more information, see Configure the user and group that run components
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-greengrass-core-v2.html#configure-component-user).
	PosixUser *string

	// The system resource limits to apply to this component's process on the core
	// device. IoT Greengrass currently supports this feature on only Linux core
	// devices. If you omit this parameter, the IoT Greengrass Core software uses the
	// default system resource limits that you configure on the Greengrass nucleus
	// component. For more information, see Configure system resource limits for
	// components
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-greengrass-core-v2.html#configure-component-system-resource-limits).
	SystemResourceLimits *SystemResourceLimits

	// The Windows user to use to run this component on Windows core devices. The user
	// must exist on each Windows core device, and its name and password must be in the
	// LocalSystem account's Credentials Manager instance. If you omit this parameter,
	// the IoT Greengrass Core software uses the default Windows user that you
	// configure on the Greengrass nucleus component. For more information, see
	// Configure the user and group that run components
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-greengrass-core-v2.html#configure-component-user).
	WindowsUser *string

	noSmithyDocumentSerde
}

// Contains information about a component version in a list.
type ComponentVersionListItem struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the component version.
	Arn *string

	// The name of the component.
	ComponentName *string

	// The version of the component.
	ComponentVersion *string

	noSmithyDocumentSerde
}

// Contains information about an endpoint and port where client devices can connect
// to an MQTT broker on a Greengrass core device.
type ConnectivityInfo struct {

	// The IP address or DNS address where client devices can connect to an MQTT broker
	// on the Greengrass core device.
	HostAddress *string

	// An ID for the connectivity information.
	Id *string

	// Additional metadata to provide to client devices that connect to this core
	// device.
	Metadata *string

	// The port where the MQTT broker operates on the core device. This port is
	// typically 8883, which is the default port for the MQTT broker component that
	// runs on core devices.
	PortNumber int32

	noSmithyDocumentSerde
}

// Contains information about a Greengrass core device, which is an IoT thing that
// runs the IoT Greengrass Core software.
type CoreDevice struct {

	// The name of the core device. This is also the name of the IoT thing.
	CoreDeviceThingName *string

	// The time at which the core device's status last updated, expressed in ISO 8601
	// format.
	LastStatusUpdateTimestamp *time.Time

	// The status of the core device. Core devices can have the following statuses:
	//
	// *
	// HEALTHY – The IoT Greengrass Core software and all components run on the core
	// device without issue.
	//
	// * UNHEALTHY – The IoT Greengrass Core software or a
	// component is in a failed state on the core device.
	Status CoreDeviceStatus

	noSmithyDocumentSerde
}

// Contains information about a deployment.
type Deployment struct {

	// The time at which the deployment was created, expressed in ISO 8601 format.
	CreationTimestamp *time.Time

	// The ID of the deployment.
	DeploymentId *string

	// The name of the deployment.
	DeploymentName *string

	// The status of the deployment.
	DeploymentStatus DeploymentStatus

	// Whether or not the deployment is the latest revision for its target.
	IsLatestForTarget bool

	// The revision number of the deployment.
	RevisionId *string

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the target IoT thing or thing group.
	TargetArn *string

	noSmithyDocumentSerde
}

// Contains information about a deployment's policy that defines when components
// are safe to update. Each component on a device can report whether or not it's
// ready to update. After a component and its dependencies are ready, they can
// apply the update in the deployment. You can configure whether or not the
// deployment notifies components of an update and waits for a response. You
// specify the amount of time each component has to respond to the update
// notification.
type DeploymentComponentUpdatePolicy struct {

	// Whether or not to notify components and wait for components to become safe to
	// update. Choose from the following options:
	//
	// * NOTIFY_COMPONENTS – The deployment
	// notifies each component before it stops and updates that component. Components
	// can use the SubscribeToComponentUpdates
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/interprocess-communication.html#ipc-operation-subscribetocomponentupdates)
	// IPC operation to receive these notifications. Then, components can respond with
	// the DeferComponentUpdate
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/interprocess-communication.html#ipc-operation-defercomponentupdate)
	// IPC operation. For more information, see Create deployments
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/create-deployments.html)
	// in the IoT Greengrass V2 Developer Guide.
	//
	// * SKIP_NOTIFY_COMPONENTS – The
	// deployment doesn't notify components or wait for them to be safe to
	// update.
	//
	// Default: NOTIFY_COMPONENTS
	Action DeploymentComponentUpdatePolicyAction

	// The amount of time in seconds that each component on a device has to report that
	// it's safe to update. If the component waits for longer than this timeout, then
	// the deployment proceeds on the device. Default: 60
	TimeoutInSeconds int32

	noSmithyDocumentSerde
}

// Contains information about how long a component on a core device can validate
// its configuration updates before it times out. Components can use the
// SubscribeToValidateConfigurationUpdates
// (https://docs.aws.amazon.com/greengrass/v2/developerguide/interprocess-communication.html#ipc-operation-subscribetovalidateconfigurationupdates)
// IPC operation to receive notifications when a deployment specifies a
// configuration update. Then, components can respond with the
// SendConfigurationValidityReport
// (https://docs.aws.amazon.com/greengrass/v2/developerguide/interprocess-communication.html#ipc-operation-sendconfigurationvalidityreport)
// IPC operation. For more information, see Create deployments
// (https://docs.aws.amazon.com/greengrass/v2/developerguide/create-deployments.html)
// in the IoT Greengrass V2 Developer Guide.
type DeploymentConfigurationValidationPolicy struct {

	// The amount of time in seconds that a component can validate its configuration
	// updates. If the validation time exceeds this timeout, then the deployment
	// proceeds for the device. Default: 30
	TimeoutInSeconds int32

	noSmithyDocumentSerde
}

// Contains information about an IoT job configuration.
type DeploymentIoTJobConfiguration struct {

	// The stop configuration for the job. This configuration defines when and how to
	// stop a job rollout.
	AbortConfig *IoTJobAbortConfig

	// The rollout configuration for the job. This configuration defines the rate at
	// which the job rolls out to the fleet of target devices.
	JobExecutionsRolloutConfig *IoTJobExecutionsRolloutConfig

	// The timeout configuration for the job. This configuration defines the amount of
	// time each device has to complete the job.
	TimeoutConfig *IoTJobTimeoutConfig

	noSmithyDocumentSerde
}

// Contains information about policies that define how a deployment updates
// components and handles failure.
type DeploymentPolicies struct {

	// The component update policy for the configuration deployment. This policy
	// defines when it's safe to deploy the configuration to devices.
	ComponentUpdatePolicy *DeploymentComponentUpdatePolicy

	// The configuration validation policy for the configuration deployment. This
	// policy defines how long each component has to validate its configure updates.
	ConfigurationValidationPolicy *DeploymentConfigurationValidationPolicy

	// The failure handling policy for the configuration deployment. This policy
	// defines what to do if the deployment fails. Default: ROLLBACK
	FailureHandlingPolicy DeploymentFailureHandlingPolicy

	noSmithyDocumentSerde
}

// Contains a request to disassociate a client device from a core device. The
// BatchDisassociateClientDeviceWithCoreDevice
// (https://docs.aws.amazon.com/greengrass/v2/APIReference/API_BatchDisassociateClientDeviceWithCoreDevice.html)
// operation consumes a list of these requests.
type DisassociateClientDeviceFromCoreDeviceEntry struct {

	// The name of the IoT thing that represents the client device to disassociate.
	//
	// This member is required.
	ThingName *string

	noSmithyDocumentSerde
}

// Contains an error that occurs from a request to disassociate a client device
// from a core device. The BatchDisassociateClientDeviceWithCoreDevice
// (https://docs.aws.amazon.com/greengrass/v2/APIReference/API_BatchDisassociateClientDeviceWithCoreDevice.html)
// operation returns a list of these errors.
type DisassociateClientDeviceFromCoreDeviceErrorEntry struct {

	// The error code for the request.
	Code *string

	// A message that provides additional information about the error.
	Message *string

	// The name of the IoT thing whose disassociate request failed.
	ThingName *string

	noSmithyDocumentSerde
}

// Contains information about a deployment job that IoT Greengrass sends to a
// Greengrass core device.
type EffectiveDeployment struct {

	// The status of the deployment job on the Greengrass core device.
	//
	// This member is required.
	CoreDeviceExecutionStatus EffectiveDeploymentExecutionStatus

	// The time at which the deployment was created, expressed in ISO 8601 format.
	//
	// This member is required.
	CreationTimestamp *time.Time

	// The ID of the deployment.
	//
	// This member is required.
	DeploymentId *string

	// The name of the deployment.
	//
	// This member is required.
	DeploymentName *string

	// The time at which the deployment job was last modified, expressed in ISO 8601
	// format.
	//
	// This member is required.
	ModifiedTimestamp *time.Time

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the target IoT thing or thing group.
	//
	// This member is required.
	TargetArn *string

	// The description of the deployment job.
	Description *string

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the IoT job that applies the deployment to target devices.
	IotJobArn *string

	// The ID of the IoT job that applies the deployment to target devices.
	IotJobId *string

	// The reason code for the update, if the job was updated.
	Reason *string

	noSmithyDocumentSerde
}

// Contains information about a component on a Greengrass core device.
type InstalledComponent struct {

	// The name of the component.
	ComponentName *string

	// The version of the component.
	ComponentVersion *string

	// Whether or not the component is a root component.
	IsRoot bool

	// The lifecycle state of the component.
	LifecycleState InstalledComponentLifecycleState

	// The details about the lifecycle state of the component.
	LifecycleStateDetails *string

	noSmithyDocumentSerde
}

// Contains a list of criteria that define when and how to cancel a configuration
// deployment.
type IoTJobAbortConfig struct {

	// The list of criteria that define when and how to cancel the configuration
	// deployment.
	//
	// This member is required.
	CriteriaList []IoTJobAbortCriteria

	noSmithyDocumentSerde
}

// Contains criteria that define when and how to cancel a job. The deployment stops
// if the following conditions are true:
//
// * The number of things that receive the
// deployment exceeds the minNumberOfExecutedThings.
//
// * The percentage of failures
// with type failureType exceeds the thresholdPercentage.
type IoTJobAbortCriteria struct {

	// The action to perform when the criteria are met.
	//
	// This member is required.
	Action IoTJobAbortAction

	// The type of job deployment failure that can cancel a job.
	//
	// This member is required.
	FailureType IoTJobExecutionFailureType

	// The minimum number of things that receive the configuration before the job can
	// cancel.
	//
	// This member is required.
	MinNumberOfExecutedThings int32

	// The minimum percentage of failureType failures that occur before the job can
	// cancel. This parameter supports up to two digits after the decimal (for example,
	// you can specify 10.9 or 10.99, but not 10.999).
	//
	// This member is required.
	ThresholdPercentage float64

	noSmithyDocumentSerde
}

// Contains information about the rollout configuration for a job. This
// configuration defines the rate at which the job deploys a configuration to a
// fleet of target devices.
type IoTJobExecutionsRolloutConfig struct {

	// The exponential rate to increase the job rollout rate.
	ExponentialRate *IoTJobExponentialRolloutRate

	// The maximum number of devices that receive a pending job notification, per
	// minute.
	MaximumPerMinute int32

	noSmithyDocumentSerde
}

// Contains information about an exponential rollout rate for a configuration
// deployment job.
type IoTJobExponentialRolloutRate struct {

	// The minimum number of devices that receive a pending job notification, per
	// minute, when the job starts. This parameter defines the initial rollout rate of
	// the job.
	//
	// This member is required.
	BaseRatePerMinute int32

	// The exponential factor to increase the rollout rate for the job. This parameter
	// supports up to one digit after the decimal (for example, you can specify 1.5,
	// but not 1.55).
	//
	// This member is required.
	IncrementFactor float64

	// The criteria to increase the rollout rate for the job.
	//
	// This member is required.
	RateIncreaseCriteria *IoTJobRateIncreaseCriteria

	noSmithyDocumentSerde
}

// Contains information about criteria to meet before a job increases its rollout
// rate. Specify either numberOfNotifiedThings or numberOfSucceededThings.
type IoTJobRateIncreaseCriteria struct {

	// The number of devices to receive the job notification before the rollout rate
	// increases.
	NumberOfNotifiedThings int32

	// The number of devices to successfully run the configuration job before the
	// rollout rate increases.
	NumberOfSucceededThings int32

	noSmithyDocumentSerde
}

// Contains information about the timeout configuration for a job.
type IoTJobTimeoutConfig struct {

	// The amount of time, in minutes, that devices have to complete the job. The timer
	// starts when the job status is set to IN_PROGRESS. If the job status doesn't
	// change to a terminal state before the time expires, then the job status is set
	// to TIMED_OUT. The timeout interval must be between 1 minute and 7 days (10080
	// minutes).
	InProgressTimeoutInMinutes int64

	noSmithyDocumentSerde
}

// Contains information about a container in which Lambda functions run on
// Greengrass core devices.
type LambdaContainerParams struct {

	// The list of system devices that the container can access.
	Devices []LambdaDeviceMount

	// The memory size of the container, expressed in kilobytes. Default: 16384 (16 MB)
	MemorySizeInKB int32

	// Whether or not the container can read information from the device's /sys folder.
	// Default: false
	MountROSysfs bool

	// The list of volumes that the container can access.
	Volumes []LambdaVolumeMount

	noSmithyDocumentSerde
}

// Contains information about a device that Linux processes in a container can
// access.
type LambdaDeviceMount struct {

	// The mount path for the device in the file system.
	//
	// This member is required.
	Path *string

	// Whether or not to add the component's system user as an owner of the device.
	// Default: false
	AddGroupOwner bool

	// The permission to access the device: read/only (ro) or read/write (rw). Default:
	// ro
	Permission LambdaFilesystemPermission

	noSmithyDocumentSerde
}

// Contains information about an event source for an Lambda function. The event
// source defines the topics on which this Lambda function subscribes to receive
// messages that run the function.
type LambdaEventSource struct {

	// The topic to which to subscribe to receive event messages.
	//
	// This member is required.
	Topic *string

	// The type of event source. Choose from the following options:
	//
	// * PUB_SUB –
	// Subscribe to local publish/subscribe messages. This event source type doesn't
	// support MQTT wildcards (+ and #) in the event source topic.
	//
	// * IOT_CORE –
	// Subscribe to Amazon Web Services IoT Core MQTT messages. This event source type
	// supports MQTT wildcards (+ and #) in the event source topic.
	//
	// This member is required.
	Type LambdaEventSourceType

	noSmithyDocumentSerde
}

// Contains parameters for a Lambda function that runs on IoT Greengrass.
type LambdaExecutionParameters struct {

	// The map of environment variables that are available to the Lambda function when
	// it runs.
	EnvironmentVariables map[string]string

	// The list of event sources to which to subscribe to receive work messages. The
	// Lambda function runs when it receives a message from an event source. You can
	// subscribe this function to local publish/subscribe messages and Amazon Web
	// Services IoT Core MQTT messages.
	EventSources []LambdaEventSource

	// The list of arguments to pass to the Lambda function when it runs.
	ExecArgs []string

	// The encoding type that the Lambda function supports. Default: json
	InputPayloadEncodingType LambdaInputPayloadEncodingType

	// The parameters for the Linux process that contains the Lambda function.
	LinuxProcessParams *LambdaLinuxProcessParams

	// The maximum amount of time in seconds that a non-pinned Lambda function can idle
	// before the IoT Greengrass Core software stops its process.
	MaxIdleTimeInSeconds int32

	// The maximum number of instances that a non-pinned Lambda function can run at the
	// same time.
	MaxInstancesCount int32

	// The maximum size of the message queue for the Lambda function component. The IoT
	// Greengrass core stores messages in a FIFO (first-in-first-out) queue until it
	// can run the Lambda function to consume each message.
	MaxQueueSize int32

	// Whether or not the Lambda function is pinned, or long-lived.
	//
	// * A pinned Lambda
	// function starts when IoT Greengrass starts and keeps running in its own
	// container.
	//
	// * A non-pinned Lambda function starts only when it receives a work
	// item and exists after it idles for maxIdleTimeInSeconds. If the function has
	// multiple work items, the IoT Greengrass Core software creates multiple instances
	// of the function.
	//
	// Default: true
	Pinned bool

	// The interval in seconds at which a pinned (also known as long-lived) Lambda
	// function component sends status updates to the Lambda manager component.
	StatusTimeoutInSeconds int32

	// The maximum amount of time in seconds that the Lambda function can process a
	// work item.
	TimeoutInSeconds int32

	noSmithyDocumentSerde
}

// Contains information about an Lambda function to import to create a component.
type LambdaFunctionRecipeSource struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the Lambda function. The ARN must include the version of the function to import.
	// You can't use version aliases like $LATEST.
	//
	// This member is required.
	LambdaArn *string

	// The component versions on which this Lambda function component depends.
	ComponentDependencies map[string]ComponentDependencyRequirement

	// The system and runtime parameters for the Lambda function as it runs on the
	// Greengrass core device.
	ComponentLambdaParameters *LambdaExecutionParameters

	// The name of the component. Defaults to the name of the Lambda function.
	ComponentName *string

	// The platforms that the component version supports.
	ComponentPlatforms []ComponentPlatform

	// The version of the component. Defaults to the version of the Lambda function as
	// a semantic version. For example, if your function version is 3, the component
	// version becomes 3.0.0.
	ComponentVersion *string

	noSmithyDocumentSerde
}

// Contains parameters for a Linux process that contains an Lambda function.
type LambdaLinuxProcessParams struct {

	// The parameters for the container in which the Lambda function runs.
	ContainerParams *LambdaContainerParams

	// The isolation mode for the process that contains the Lambda function. The
	// process can run in an isolated runtime environment inside the IoT Greengrass
	// container, or as a regular process outside any container. Default:
	// GreengrassContainer
	IsolationMode LambdaIsolationMode

	noSmithyDocumentSerde
}

// Contains information about a volume that Linux processes in a container can
// access. When you define a volume, the IoT Greengrass Core software mounts the
// source files to the destination inside the container.
type LambdaVolumeMount struct {

	// The path to the logical volume in the file system.
	//
	// This member is required.
	DestinationPath *string

	// The path to the physical volume in the file system.
	//
	// This member is required.
	SourcePath *string

	// Whether or not to add the IoT Greengrass user group as an owner of the volume.
	// Default: false
	AddGroupOwner bool

	// The permission to access the volume: read/only (ro) or read/write (rw). Default:
	// ro
	Permission LambdaFilesystemPermission

	noSmithyDocumentSerde
}

// Contains information about a component version that is compatible to run on a
// Greengrass core device.
type ResolvedComponentVersion struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the component version.
	Arn *string

	// The name of the component.
	ComponentName *string

	// The version of the component.
	ComponentVersion *string

	// A message that communicates details about the vendor guidance state of the
	// component version. This message communicates why a component version is
	// discontinued or deleted.
	Message *string

	// The recipe of the component version.
	Recipe []byte

	// The vendor guidance state for the component version. This state indicates
	// whether the component version has any issues that you should consider before you
	// deploy it. The vendor guidance state can be:
	//
	// * ACTIVE – This component version
	// is available and recommended for use.
	//
	// * DISCONTINUED – This component version
	// has been discontinued by its publisher. You can deploy this component version,
	// but we recommend that you use a different version of this component.
	//
	// * DELETED
	// – This component version has been deleted by its publisher, so you can't deploy
	// it. If you have any existing deployments that specify this component version,
	// those deployments will fail.
	VendorGuidance VendorGuidance

	noSmithyDocumentSerde
}

// Contains information about system resource limits that the IoT Greengrass Core
// software applies to a component's processes. For more information, see Configure
// system resource limits for components
// (https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-greengrass-core-v2.html#configure-component-system-resource-limits).
type SystemResourceLimits struct {

	// The maximum amount of CPU time that a component's processes can use on the core
	// device. A core device's total CPU time is equivalent to the device's number of
	// CPU cores. For example, on a core device with 4 CPU cores, you can set this
	// value to 2 to limit the component's processes to 50 percent usage of each CPU
	// core. On a device with 1 CPU core, you can set this value to 0.25 to limit the
	// component's processes to 25 percent usage of the CPU. If you set this value to a
	// number greater than the number of CPU cores, the IoT Greengrass Core software
	// doesn't limit the component's CPU usage.
	Cpus float64

	// The maximum amount of RAM, expressed in kilobytes, that a component's processes
	// can use on the core device.
	Memory int64

	noSmithyDocumentSerde
}

// Contains information about a validation exception field.
type ValidationExceptionField struct {

	// The message of the exception field.
	//
	// This member is required.
	Message *string

	// The name of the exception field.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
