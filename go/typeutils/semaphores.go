package typeutils

import (
	"context"
	"errors"
	"golang.org/x/sync/semaphore"
	"log"
	"time"
)

type Semaphores struct {
	semaphoreMap map[string]Semaphore
}

type Semaphore struct {
	name     string
	maximum  int
	value    int
	weighted *semaphore.Weighted
	timeout  int
}

func NewSemaphores() Semaphores {
	return Semaphores{
		semaphoreMap: make(map[string]Semaphore),
	}
}

func (s *Semaphores) Add(name string, maximum int, timeout int) {
	s.semaphoreMap[name] = Semaphore{
		name:     name,
		maximum:  maximum,
		value:    maximum,
		weighted: semaphore.NewWeighted(int64(maximum)),
		timeout:  timeout,
	}
}

func (s *Semaphores) P(name string) error {
	return s.Wait(name)
}

func (s *Semaphores) Wait(name string) error {
	if sem, defined := s.semaphoreMap[name]; defined {
		sem.value--
		s.semaphoreMap[name] = sem

		startTime := time.Now()
		timeout := sem.timeout
		for s.semaphoreMap[name].value < 0 {
			nowTime := time.Now()
			if nowTime.Second()%30 == 0 {
				log.Println("waiting on semaphore:", name)
			}
			if nowTime.After(startTime.Add(time.Duration(timeout) * time.Second)) {
				sem.value++
				if sem.value > sem.maximum {
					sem.value = sem.maximum
				}
				s.semaphoreMap[name] = sem
				return errors.New("timeout for wait on semaphore: " + name)
			}
		}
		return nil
	}

	return errors.New("invalid semaphore name: " + name)
}

func (s *Semaphores) Acquire(name string) error {
	if sem, defined := s.semaphoreMap[name]; defined {
		timeout := time.Now().Add(time.Duration(sem.timeout) * time.Second)
		ctx, cancel := context.WithDeadline(context.Background(), timeout)
		defer cancel()

		err := sem.weighted.Acquire(ctx, 1)
		if err != nil {
			return errors.New("timeout for wait on semaphore: " + name + ": " + ctx.Err().Error())
		}

		return nil
	}

	return errors.New("invalid semaphore name: " + name)
}

func (s *Semaphores) V(name string) error {
	return s.Signal(name)
}

func (s *Semaphores) Signal(name string) error {
	if sem, defined := s.semaphoreMap[name]; defined {
		sem.value++
		if sem.value > sem.maximum {
			sem.value = sem.maximum
		}
		s.semaphoreMap[name] = sem
		return nil
	}

	return errors.New("invalid semaphore name: " + name)
}

func (s *Semaphores) Release(name string) error {
	if sem, defined := s.semaphoreMap[name]; defined {
		sem.weighted.Release(1)
		return nil
	}

	return errors.New("invalid semaphore name: " + name)
}
