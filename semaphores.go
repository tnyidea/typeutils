package typeutils

import (
	"errors"
	"log"
	"time"
)

type Semaphores struct {
	semaphoreMap map[string]Semaphore
}

type Semaphore struct {
	Name         string
	MaximumValue int
	value        int
	Timeout      int
}

func (s *Semaphores) Add(semaphore Semaphore) {
	s.semaphoreMap[semaphore.Name] = semaphore
}

func (s *Semaphores) P(name string) error {
	return s.Wait(name)
}

func (s *Semaphores) Wait(name string) error {
	if semaphore, defined := s.semaphoreMap[name]; defined {
		semaphore.value--
		s.semaphoreMap[name] = semaphore

		startTime := time.Now()
		timeout := semaphore.Timeout
		for s.semaphoreMap[name].value < 0 {
			log.Println("wait on semaphore:", name)
			nowTime := time.Now()
			if nowTime.After(startTime.Add(time.Duration(timeout) * time.Second)) {
				return errors.New("timeout for wait on semaphore: " + name)
			}
		}
		return nil
	}

	return errors.New("invalid semaphore name: " + name)
}

func (s *Semaphores) V(name string) error {
	return s.Signal(name)
}

func (s *Semaphores) Signal(name string) error {
	if semaphore, defined := s.semaphoreMap[name]; defined {
		semaphore.value++
		if semaphore.value > semaphore.MaximumValue {
			semaphore.value = semaphore.MaximumValue
		}
		s.semaphoreMap[name] = semaphore
		return nil
	}

	return errors.New("invalid semaphore name: " + name)
}
