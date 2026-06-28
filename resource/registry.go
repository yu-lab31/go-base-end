// Resource like database, cache or logger registry their providers here.
// Hence, this package is depended on by (almost) all other packages.
// Providers create resources or return singleton instances; it is recommended for resource like database
// link pools to use singleton because the cost of creating new instances is too heavy.
package resource

import (
	"fmt"
	"reflect"
	"sync"
)

// Option is the interface that wraps methods setting up resource.
type Option interface {
	Set(any)
}

type Provider func(opts ...Option) any

// registry contains resources' names and providers.
type registry struct {
	factories map[string]Provider
	mtx       *sync.RWMutex
}

var (
	reg                registry
	DuplicatedResource = fmt.Errorf("resource with such type has already been registered")
	ResourceNotExist   = fmt.Errorf("resource with such type hasn't been registered")
)

// Register resource with its type name and provider.
func Register[T any](provider Provider) error {
	reg.mtx.Lock()
	defer reg.mtx.Unlock()

	name := getResourceName[T]()

	if _, ok := reg.factories[name]; ok {
		return fmt.Errorf("%w: %s", DuplicatedResource, name)
	}
	reg.factories[name] = provider

	return nil
}

// Get returns resource with type T; returns ResourceNotExist error when resource with type T is
// not registered or failed in type check (like T=Resource but registered *Resource).
// Be aware that it's up to the resource package whether `opts` could behave or not (for example,
// some settings may not be changed after the resources' creation).
func Get[T any](opts ...Option) (T, error) {
	reg.mtx.RLock()
	defer reg.mtx.RUnlock()

	var zero, resource T

	name := getResourceName[T]()

	if prov, ok := reg.factories[name]; !ok {
		return zero, fmt.Errorf("%w: %s", ResourceNotExist, name)
	} else if resource, ok = prov(opts...).(T); !ok {
		return zero, fmt.Errorf("%w (wrong type): %s", ResourceNotExist, name)
	}

	return resource, nil
}

func getResourceName[T any]() string {
	t := reflect.TypeFor[T]()

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	return t.PkgPath() + t.Name()
}

func init() {
	reg = registry{
		factories: make(map[string]Provider),
		mtx:       new(sync.RWMutex),
	}
}
