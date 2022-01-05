package handler

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/netw-device-driver/ndd-runtime/pkg/utils"
	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/logging"
	asv1alpha1 "github.com/yndd/nddr-as-registry/apis/as/v1alpha1"
	"github.com/yndd/nddr-as-registry/internal/pool"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func New(opts ...Option) (Handler, error) {
	rgfn := func() asv1alpha1.Rg { return &asv1alpha1.Registry{} }
	s := &handler{
		pool:        make(map[string]pool.Pool),
		speedy:      make(map[string]int),
		newRegistry: rgfn,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s, nil
}

func (r *handler) WithLogger(log logging.Logger) {
	r.log = log
}

func (r *handler) WithClient(c client.Client) {
	r.client = c
}

type RegisterInfo struct {
	Namespace    string
	Name         string
	RegistryName string
	CrName       string
	Index        uint32
	Selector     map[string]string
	SourceTag    map[string]string
}

type handler struct {
	log logging.Logger
	// kubernetes
	client client.Client

	newRegistry func() asv1alpha1.Rg
	poolMutex   sync.Mutex
	pool        map[string]pool.Pool
	speedyMutex sync.Mutex
	speedy      map[string]int
}

func (r *handler) Init(crName string, start, end uint32, allocStrategy string) {
	r.poolMutex.Lock()
	defer r.poolMutex.Unlock()
	if _, ok := r.pool[crName]; !ok {
		r.pool[crName] = pool.New(start, end, allocStrategy)
	}

	r.speedyMutex.Lock()
	defer r.speedyMutex.Unlock()
	if _, ok := r.speedy[crName]; !ok {
		r.speedy[crName] = 0
	}
}

func (r *handler) Delete(crName string) {
	r.poolMutex.Lock()
	defer r.poolMutex.Unlock()
	delete(r.pool, crName)

	r.speedyMutex.Lock()
	defer r.speedyMutex.Unlock()
	delete(r.speedy, crName)
}

func (r *handler) GetAllocated(crName string) (uint32, []*uint32) {
	r.poolMutex.Lock()
	defer r.poolMutex.Unlock()
	if pool, ok := r.pool[crName]; ok {
		return pool.GetAllocated()
	}
	return 0, make([]*uint32, 0)
}

func (r *handler) ResetSpeedy(crName string) {
	r.speedyMutex.Lock()
	defer r.speedyMutex.Unlock()
	if _, ok := r.speedy[crName]; ok {
		r.speedy[crName] = 0
	}
}

func (r *handler) GetSpeedy(crName string) int {
	r.speedyMutex.Lock()
	defer r.speedyMutex.Unlock()
	if _, ok := r.speedy[crName]; ok {
		return r.speedy[crName]
	}
	return 9999
}

func (r *handler) IncrementSpeedy(crName string) {
	r.speedyMutex.Lock()
	defer r.speedyMutex.Unlock()
	if _, ok := r.speedy[crName]; ok {
		r.speedy[crName]++
	}
}

func (r *handler) Register(ctx context.Context, info *RegisterInfo) (*uint32, error) {
	pool, index, err := r.validateRegister(ctx, info)
	if err != nil {
		return nil, err
	}
	requestName := info.Name
	sourceTag := info.SourceTag

	r.log.Debug("pool insert", "index", index)
	as := pool.InsertByIndex(*index, requestName, sourceTag)
	r.log.Debug("pool inserted", "index", index, "as", as)

	return &as, nil
}

func (r *handler) DeRegister(ctx context.Context, info *RegisterInfo) error {

	pool, index, err := r.validateRegister(ctx, info)
	if err != nil {
		return err
	}
	requestName := info.Name
	sourceTag := info.SourceTag

	r.log.Debug("pool delete", "index", index)
	pool.DeleteByIndex(*index, requestName, sourceTag)
	r.log.Debug("pool deleted", "index", index)

	return nil
}

func (r *handler) validateRegister(ctx context.Context, info *RegisterInfo) (pool.Pool, *uint32, error) {
	namespace := info.Namespace
	registryName := info.RegistryName
	crName := info.CrName
	selector := info.Selector

	// find registry in k8s api
	registry := r.newRegistry()
	if err := r.client.Get(ctx, types.NamespacedName{
		Namespace: namespace,
		Name:      registryName}, registry); err != nil {
		// can happen when the ipam is not found
		r.log.Debug("registry not found")
		return nil, nil, errors.Wrap(err, "registry not found")
	}

	// check is registry is ready
	if registry.GetCondition(asv1alpha1.ConditionKindReady).Status != corev1.ConditionTrue {
		r.log.Debug("Registry not ready")
		return nil, nil, errors.New("Registry not ready")
	}

	// check if the supplied info is available
	if _, ok := selector["index"]; !ok {
		return nil, nil, errors.New("selector does not contain a index")
	}
	index := selector["index"]
	idx, err := strconv.Atoi(index)
	if err != nil {
		return nil, nil, err
	}

	// check if the pool/register is ready to handle new registrations
	r.poolMutex.Lock()
	defer r.poolMutex.Unlock()
	if _, ok := r.pool[crName]; !ok {
		r.log.Debug("pool/tree not ready", "crName", crName)
		return nil, nil, fmt.Errorf("pool/tree not ready, crName: %s", crName)
	}
	pool := r.pool[crName]

	return pool, utils.Uint32Ptr(uint32(idx)), nil
}
