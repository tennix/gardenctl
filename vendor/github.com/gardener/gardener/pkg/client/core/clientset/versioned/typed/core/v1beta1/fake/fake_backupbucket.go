// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackupBuckets implements BackupBucketInterface
type FakeBackupBuckets struct {
	Fake *FakeCoreV1beta1
}

var backupbucketsResource = schema.GroupVersionResource{Group: "core.gardener.cloud", Version: "v1beta1", Resource: "backupbuckets"}

var backupbucketsKind = schema.GroupVersionKind{Group: "core.gardener.cloud", Version: "v1beta1", Kind: "BackupBucket"}

// Get takes name of the backupBucket, and returns the corresponding backupBucket object, and an error if there is any.
func (c *FakeBackupBuckets) Get(name string, options v1.GetOptions) (result *v1beta1.BackupBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(backupbucketsResource, name), &v1beta1.BackupBucket{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupBucket), err
}

// List takes label and field selectors, and returns the list of BackupBuckets that match those selectors.
func (c *FakeBackupBuckets) List(opts v1.ListOptions) (result *v1beta1.BackupBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(backupbucketsResource, backupbucketsKind, opts), &v1beta1.BackupBucketList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BackupBucketList{ListMeta: obj.(*v1beta1.BackupBucketList).ListMeta}
	for _, item := range obj.(*v1beta1.BackupBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupBuckets.
func (c *FakeBackupBuckets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(backupbucketsResource, opts))
}

// Create takes the representation of a backupBucket and creates it.  Returns the server's representation of the backupBucket, and an error, if there is any.
func (c *FakeBackupBuckets) Create(backupBucket *v1beta1.BackupBucket) (result *v1beta1.BackupBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(backupbucketsResource, backupBucket), &v1beta1.BackupBucket{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupBucket), err
}

// Update takes the representation of a backupBucket and updates it. Returns the server's representation of the backupBucket, and an error, if there is any.
func (c *FakeBackupBuckets) Update(backupBucket *v1beta1.BackupBucket) (result *v1beta1.BackupBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(backupbucketsResource, backupBucket), &v1beta1.BackupBucket{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupBucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupBuckets) UpdateStatus(backupBucket *v1beta1.BackupBucket) (*v1beta1.BackupBucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(backupbucketsResource, "status", backupBucket), &v1beta1.BackupBucket{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupBucket), err
}

// Delete takes name of the backupBucket and deletes it. Returns an error if one occurs.
func (c *FakeBackupBuckets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(backupbucketsResource, name), &v1beta1.BackupBucket{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupBuckets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(backupbucketsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.BackupBucketList{})
	return err
}

// Patch applies the patch and returns the patched backupBucket.
func (c *FakeBackupBuckets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.BackupBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(backupbucketsResource, name, pt, data, subresources...), &v1beta1.BackupBucket{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupBucket), err
}
