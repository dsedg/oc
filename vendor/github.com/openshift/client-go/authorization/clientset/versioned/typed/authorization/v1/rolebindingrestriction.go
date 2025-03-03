// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	authorizationv1 "github.com/openshift/api/authorization/v1"
	applyconfigurationsauthorizationv1 "github.com/openshift/client-go/authorization/applyconfigurations/authorization/v1"
	scheme "github.com/openshift/client-go/authorization/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// RoleBindingRestrictionsGetter has a method to return a RoleBindingRestrictionInterface.
// A group's client should implement this interface.
type RoleBindingRestrictionsGetter interface {
	RoleBindingRestrictions(namespace string) RoleBindingRestrictionInterface
}

// RoleBindingRestrictionInterface has methods to work with RoleBindingRestriction resources.
type RoleBindingRestrictionInterface interface {
	Create(ctx context.Context, roleBindingRestriction *authorizationv1.RoleBindingRestriction, opts metav1.CreateOptions) (*authorizationv1.RoleBindingRestriction, error)
	Update(ctx context.Context, roleBindingRestriction *authorizationv1.RoleBindingRestriction, opts metav1.UpdateOptions) (*authorizationv1.RoleBindingRestriction, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*authorizationv1.RoleBindingRestriction, error)
	List(ctx context.Context, opts metav1.ListOptions) (*authorizationv1.RoleBindingRestrictionList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *authorizationv1.RoleBindingRestriction, err error)
	Apply(ctx context.Context, roleBindingRestriction *applyconfigurationsauthorizationv1.RoleBindingRestrictionApplyConfiguration, opts metav1.ApplyOptions) (result *authorizationv1.RoleBindingRestriction, err error)
	RoleBindingRestrictionExpansion
}

// roleBindingRestrictions implements RoleBindingRestrictionInterface
type roleBindingRestrictions struct {
	*gentype.ClientWithListAndApply[*authorizationv1.RoleBindingRestriction, *authorizationv1.RoleBindingRestrictionList, *applyconfigurationsauthorizationv1.RoleBindingRestrictionApplyConfiguration]
}

// newRoleBindingRestrictions returns a RoleBindingRestrictions
func newRoleBindingRestrictions(c *AuthorizationV1Client, namespace string) *roleBindingRestrictions {
	return &roleBindingRestrictions{
		gentype.NewClientWithListAndApply[*authorizationv1.RoleBindingRestriction, *authorizationv1.RoleBindingRestrictionList, *applyconfigurationsauthorizationv1.RoleBindingRestrictionApplyConfiguration](
			"rolebindingrestrictions",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *authorizationv1.RoleBindingRestriction { return &authorizationv1.RoleBindingRestriction{} },
			func() *authorizationv1.RoleBindingRestrictionList {
				return &authorizationv1.RoleBindingRestrictionList{}
			},
		),
	}
}
