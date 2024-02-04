package formFormula

import (
	"fmt"
	"sync"
	"time"
)

type WorkersPool interface {
	Start()
}

type Scheduler interface {
	Next() chan Job
	ReportThatJobDone(uint64)
}

type Job interface {
	GetIndex() uint64
	ToString() string
	Done()
	IsDone() bool
	IncrementCycle()
	Stat() string
}

type job struct {
	index         uint64
	serialization string
	done          bool

	startTime  time.Time
	finishTime time.Time
	cycles     uint64
}

type scheduler struct {
	startedJobsPipe      chan Job
	finishedJobsPipe     chan uint64
	headJobSerialization string
	head                 uint64
	tail                 uint64
	buffer               []Job
	jobConstructor       func(currentJob string) string
	configSaver          func(lastJob Job)
}

func NewJob(index uint64, serialization string) Job {
	return &job{
		index:         index,
		serialization: serialization,
		startTime:     time.Now(),
		finishTime:    time.Now(),
		cycles:        0,
	}
}

func (job *job) Done() {
	job.finishTime = time.Now()
	job.done = true
}

func (job *job) IsDone() bool {
	return job.done
}

func (job *job) IncrementCycle() {
	job.cycles++
}

func (job *job) ToString() string {
	return job.serialization
}

func (job *job) Stat() string {
	mils := job.finishTime.Sub(job.startTime).Milliseconds()
	if mils == 0 {
		mils = 1
	}
	return fmt.Sprintf("elapsed(ms):%10v [%10v cycles] - %4v cycle/ms", mils, job.cycles, job.cycles/uint64(mils))
}

func (job *job) GetIndex() uint64 {
	return job.index
}

func (scheduler *scheduler) Next() chan Job {
	return scheduler.startedJobsPipe
}

func (scheduler *scheduler) ReportThatJobDone(indexOfFinishedJob uint64) {
	//fmt.Printf("report done %v", indexOfFinishedJob)
	scheduler.finishedJobsPipe <- indexOfFinishedJob
}

func newScheduler(lastFinishedJobIndex uint64, lastFinishedJob string, doneJobsBufferSize uint, jobConstructor func(currentJob string) string, configSaver func(lastJob Job)) Scheduler {
	scheduler := &scheduler{
		startedJobsPipe:      make(chan Job, doneJobsBufferSize+1),
		finishedJobsPipe:     make(chan uint64, doneJobsBufferSize+1),
		head:                 lastFinishedJobIndex + 1,
		tail:                 lastFinishedJobIndex + 1,
		headJobSerialization: lastFinishedJob,
		jobConstructor:       jobConstructor,
		buffer:               make([]Job, doneJobsBufferSize),
		configSaver:          configSaver,
	}

	go func() {
		for {
			select {
			case finishedJobIndex := <-scheduler.finishedJobsPipe:
				jobAddr := finishedJobIndex % uint64(len(scheduler.buffer))
				scheduler.buffer[jobAddr].Done()
			default:
				{
					// clear closed jobs from buffer
					if scheduler.tail <= scheduler.head {
						tailAddr := scheduler.tail % uint64(len(scheduler.buffer))
						job := scheduler.buffer[tailAddr]
						if job != nil {
							if job.IsDone() {
								scheduler.configSaver(job)
								scheduler.buffer[tailAddr] = nil
								scheduler.tail++
							}
						}
					}

					// start new jobs in buffer
					if scheduler.head < scheduler.tail+uint64(len(scheduler.buffer)) {
						headAddr := scheduler.head % uint64(len(scheduler.buffer))
						if scheduler.buffer[headAddr] == nil {

							scheduler.headJobSerialization = jobConstructor(scheduler.headJobSerialization)
							newJob := NewJob(scheduler.head, scheduler.headJobSerialization)
							scheduler.buffer[headAddr] = newJob
							scheduler.head++
							scheduler.startedJobsPipe <- newJob
						}
					}
				}
			}
		}
	}()

	return scheduler

}

type workersPool struct {
	wg        sync.WaitGroup
	handler   func(threadIndex uint, job Job) bool
	threads   uint
	scheduler Scheduler
}

func NewWorkersPool(lastFinishedJobIndex uint64, lastFinishedJob string, maxThreads uint, maxDoneJobsBuffer uint, handler func(threadIndex uint, job Job) bool, configSaver func(job Job)) WorkersPool {

	jobConstructor := func(currentSequence string) string {
		newSequence, err := GetNextBracketsSequence(currentSequence, 2)
		if err != nil {
			panic(err)
		}
		return newSequence
	}

	newWorkersPool := &workersPool{
		threads:   maxThreads,
		handler:   handler,
		scheduler: newScheduler(lastFinishedJobIndex, lastFinishedJob, maxDoneJobsBuffer, jobConstructor, configSaver),
	}

	return newWorkersPool
}

func (workersPool *workersPool) Start() {

	for i := uint(0); i < workersPool.threads; i++ {

		workersPool.wg.Add(1)
		go func(threadIndex uint) {

		exitLoop:
			for {
				job := <-workersPool.scheduler.Next()

				hasNext := workersPool.handler(threadIndex, job)

				workersPool.scheduler.ReportThatJobDone(job.GetIndex())

				if !hasNext {
					break exitLoop
				}

			}
			//fmt.Printf("thread %v exited\n", threadIndex)

			workersPool.wg.Done()
		}(i)
	}

	workersPool.wg.Wait()
}
