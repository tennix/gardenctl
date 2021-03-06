// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	scheme "github.com/gardener/gardener/pkg/client/core/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ShootStatesGetter has a method to return a ShootStateInterface.
// A group's client should implement this interface.
type ShootStatesGetter interface {
	ShootStates(namespace string) ShootStateInterface
}

// ShootStateInterface has methods to work with ShootState resources.
type ShootStateInterface interface {
	Create(*v1alpha1.ShootState) (*v1alpha1.ShootState, error)
	Update(*v1alpha1.ShootState) (*v1alpha1.ShootState, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ShootState, error)
	List(opts v1.ListOptions) (*v1alpha1.ShootStateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ShootState, err error)
	ShootStateExpansion
}

// shootStates implements ShootStateInterface
type shootStates struct {
	client rest.Interface
	ns     string
}

// newShootStates returns a ShootStates
func newShootStates(c *CoreV1alpha1Client, namespace string) *shootStates {
	return &shootStates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the shootState, and returns the corresponding shootState object, and an error if there is any.
func (c *shootStates) Get(name string, options v1.GetOptions) (result *v1alpha1.ShootState, err error) {
	result = &v1alpha1.ShootState{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("shootstates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ShootStates that match those selectors.
func (c *shootStates) List(opts v1.ListOptions) (result *v1alpha1.ShootStateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ShootStateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("shootstates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested shootStates.
func (c *shootStates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("shootstates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a shootState and creates it.  Returns the server's representation of the shootState, and an error, if there is any.
func (c *shootStates) Create(shootState *v1alpha1.ShootState) (result *v1alpha1.ShootState, err error) {
	result = &v1alpha1.ShootState{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("shootstates").
		Body(shootState).
		Do().
		Into(result)
	return
}

// Update takes the representation of a shootState and updates it. Returns the server's representation of the shootState, and an error, if there is any.
func (c *shootStates) Update(shootState *v1alpha1.ShootState) (result *v1alpha1.ShootState, err error) {
	result = &v1alpha1.ShootState{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("shootstates").
		Name(shootState.Name).
		Body(shootState).
		Do().
		Into(result)
	return
}

// Delete takes name of the shootState and deletes it. Returns an error if one occurs.
func (c *shootStates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("shootstates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *shootStates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("shootstates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched shootState.
func (c *shootStates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ShootState, err error) {
	result = &v1alpha1.ShootState{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("shootstates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
