// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	"github.com/gardener/gardener/pkg/client/core/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type CoreV1alpha1Interface interface {
	RESTClient() rest.Interface
	BackupBucketsGetter
	BackupEntriesGetter
	CloudProfilesGetter
	ControllerInstallationsGetter
	ControllerRegistrationsGetter
	PlantsGetter
	ProjectsGetter
	QuotasGetter
	SecretBindingsGetter
	SeedsGetter
	ShootsGetter
	ShootStatesGetter
}

// CoreV1alpha1Client is used to interact with features provided by the core.gardener.cloud group.
type CoreV1alpha1Client struct {
	restClient rest.Interface
}

func (c *CoreV1alpha1Client) BackupBuckets() BackupBucketInterface {
	return newBackupBuckets(c)
}

func (c *CoreV1alpha1Client) BackupEntries(namespace string) BackupEntryInterface {
	return newBackupEntries(c, namespace)
}

func (c *CoreV1alpha1Client) CloudProfiles() CloudProfileInterface {
	return newCloudProfiles(c)
}

func (c *CoreV1alpha1Client) ControllerInstallations() ControllerInstallationInterface {
	return newControllerInstallations(c)
}

func (c *CoreV1alpha1Client) ControllerRegistrations() ControllerRegistrationInterface {
	return newControllerRegistrations(c)
}

func (c *CoreV1alpha1Client) Plants(namespace string) PlantInterface {
	return newPlants(c, namespace)
}

func (c *CoreV1alpha1Client) Projects() ProjectInterface {
	return newProjects(c)
}

func (c *CoreV1alpha1Client) Quotas(namespace string) QuotaInterface {
	return newQuotas(c, namespace)
}

func (c *CoreV1alpha1Client) SecretBindings(namespace string) SecretBindingInterface {
	return newSecretBindings(c, namespace)
}

func (c *CoreV1alpha1Client) Seeds() SeedInterface {
	return newSeeds(c)
}

func (c *CoreV1alpha1Client) Shoots(namespace string) ShootInterface {
	return newShoots(c, namespace)
}

func (c *CoreV1alpha1Client) ShootStates(namespace string) ShootStateInterface {
	return newShootStates(c, namespace)
}

// NewForConfig creates a new CoreV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*CoreV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CoreV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new CoreV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CoreV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CoreV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *CoreV1alpha1Client {
	return &CoreV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CoreV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
