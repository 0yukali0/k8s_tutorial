package controller

import (
	"context"
	"fmt"
	"time"

	"k8s_monitor/pkg/common"

	v1 "k8s.io/api/core/v1"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	queue    workqueue.TypedRateLimitingInterface[string]
	informer cache.SharedIndexInformer 
}

func NewController(queue workqueue.TypedRateLimitingInterface[string], informer cache.SharedIndexInformer ) *Controller {
	return &Controller{
		informer: informer,
		queue:    queue,
	}
}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)
	err := c.syncToStdout(key)
	c.handleErr(err, key)
	return true
}

func (c *Controller) syncToStdout(key string) error {
	obj, exists, err := c.informer.GetIndexer().GetByKey(key)
	if err != nil {
		common.GetLogger().Error("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		common.GetLogger().Info("Pod %s does not exist anymore\n", key)
	} else {
		common.GetLogger().Info("Sync/Add/Update for Pod %s\n", obj.(*v1.Pod).GetName())
	}
	return nil
}

func (c *Controller) handleErr(err error, key string) {
	if err == nil {
		c.queue.Forget(key)
		return
	}

	if c.queue.NumRequeues(key) < 5 {
		common.GetLogger().WithFields(logrus.Fields{
			"pod": key,
			"err": err.Error(),
		}).Info("Error syncing pod")
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)
	runtime.HandleError(err)
	common.GetLogger().WithFields(logrus.Fields{
		"pod": key,
		"err": err.Error(),
	}).Info("Dropping pod")
}

func (c *Controller) Run(ctx context.Context, workers int) {
	defer runtime.HandleCrash()

	defer c.queue.ShutDown()
	common.GetLogger().Info("Starting Pod controller")

	go c.informer.Run(ctx.Done())

	if !cache.WaitForCacheSync(ctx.Done(), c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	for i := 0; i < workers; i++ {
		go wait.Until(c.runWorker, time.Second, ctx.Done())
	}

	for {
		select {
		case <-ctx.Done():
			common.GetLogger().Info("Stopping controller")
			return
		default:
			common.GetLogger().Info("Running controller")
		}
	}
}

func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}
