package timer

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type Timer interface {
	AddTaskByFunc(cronName string, spec string, taskFunc func(), taskName string, option ...cron.Option) (cron.EntryID, error)
	StopCron(cronName string)
}

type task struct {
	EntryID  cron.EntryID
	Spec     string
	TaskName string
}

type taskManager struct {
	cron  *cron.Cron
	tasks map[cron.EntryID]*task
}

type timer struct {
	cronList map[string]*taskManager
	sync.Mutex
}

func NewTimerTask() Timer {
	return &timer{
		cronList: make(map[string]*taskManager),
	}
}

// AddTaskByFunc 添加定时任务
func (t *timer) AddTaskByFunc(cronName string, spec string, taskFunc func(), taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			cron:  cron.New(option...),
			tasks: tasks,
		}
	}
	entryID, err := t.cronList[cronName].cron.AddFunc(spec, taskFunc)
	t.cronList[cronName].cron.Start()
	t.cronList[cronName].tasks[entryID] = &task{
		EntryID:  entryID,
		Spec:     spec,
		TaskName: taskName,
	}
	return entryID, err
}

// StopCron 停止任务
func (t *timer) StopCron(cronName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.cron.Stop()
	}
}
