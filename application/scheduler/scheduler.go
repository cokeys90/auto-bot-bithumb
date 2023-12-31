package scheduler

import (
	"context"
	"github.com/cokeys90/auto-bot-bithumb/application/bithumb/connector"
	"github.com/reugn/go-quartz/quartz"
	"sync"
)

var instance *Manager

type Manager struct {
	scheduler quartz.Scheduler
	mutex     *sync.Mutex
}

// initialize 스케줄러 초기화
func (m *Manager) initialize(_ctx context.Context) {
	m.mutex = new(sync.Mutex)

	m.scheduler = quartz.NewStdScheduler()
	m.scheduler.Start(_ctx)
	go func() {
		m.scheduler.Wait(_ctx)
	}()
}

// Instance 스케줄러 인스턴스 반환
func Instance() *Manager {
	if instance == nil {
		instance = new(Manager)
		instance.initialize(context.Background())
	}
	return instance
}

func (m *Manager) StartAutoBotBithumb(_connector connector.RextConnector)
