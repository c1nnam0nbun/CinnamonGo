package ecs

import (
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"sync"
)

type Stage string

const (
	First      Stage = "First"
	PreUpdate  Stage = "PreUpdate"
	Update     Stage = "Update"
	PostUpdate Stage = "PostUpdate"
	Last       Stage = "Last"
)

type scheduler struct {
	world           *World
	stages          map[Stage][]System
	stageOrder      []Stage
	startupSystems  []System
	shutdownSystems []System
}

func NewScheduler(world *World) Scheduler {
	return &scheduler{
		world:      world,
		stageOrder: []Stage{First, PreUpdate, Update, PostUpdate, Last},
		stages: map[Stage][]System{
			First:      make([]System, 0),
			PreUpdate:  make([]System, 0),
			Update:     make([]System, 0),
			PostUpdate: make([]System, 0),
			Last:       make([]System, 0),
		},
	}
}

type Scheduler interface {
	AddStage(stage Stage)
	AddStageBefore(existingStage, stage Stage)
	AddStageAfter(existingStage, stage Stage)
	AddStartupSystem(system System)
	AddShutdownSystem(system System)
	AddSystem(system System)
	AddSystemToStage(stage Stage, system System)
	Run()
}

func (s *scheduler) AddStage(stage Stage) {
	s.stageOrder = append(s.stageOrder, stage)
	s.stages[stage] = make([]System, 0)
}

func (s *scheduler) AddStageBefore(existingStage, stage Stage) {
	index := slices.Index(s.stageOrder, existingStage)
	if index == -1 {
		s.AddStage(stage)
		return
	}
	s.stageOrder = slices.Insert(s.stageOrder, index-1, stage)
	s.stages[stage] = make([]System, 0)
}

func (s *scheduler) AddStageAfter(existingStage, stage Stage) {
	index := slices.Index(s.stageOrder, existingStage)
	if index == -1 || index == len(s.stageOrder)-1 {
		s.AddStage(stage)
		return
	}
	s.stageOrder = slices.Insert(s.stageOrder, index+1, stage)
	s.stages[stage] = make([]System, 0)
}

func (s *scheduler) AddStartupSystem(system System) {
	s.startupSystems = append(s.startupSystems, system)
}

func (s *scheduler) AddShutdownSystem(system System) {
	s.shutdownSystems = append(s.shutdownSystems, system)
}

func (s *scheduler) AddSystem(system System) {
	s.stages[Update] = append(s.stages[Update], system)
}

func (s *scheduler) AddSystemToStage(stage Stage, system System) {
	if !slices.Contains(maps.Keys(s.stages), stage) {
		return
	}
	s.stages[stage] = append(s.stages[stage], system)
}

func (s *scheduler) Run() {
	if len(s.startupSystems) != 0 {
		for _, system := range s.startupSystems {
			system.call(s.world)
		}
		s.startupSystems = make([]System, 0)
	}
	for _, stage := range s.stageOrder {
		println(stage)
		var wg sync.WaitGroup
		for _, system := range s.stages[stage] {
			go func() {
				system.call(s.world)
				wg.Done()
			}()
			wg.Add(1)
		}
		wg.Wait()
	}
}
