package core

import (
	"github.com/c1nnam0nbun/cinnamon/window"
	"github.com/c1nnqm0nbun/cinnamon/ecs"
)

type Application struct {
	world  *ecs.World
	window window.Window
}

func NewApplication() *Application {
	return &Application{
		world:  ecs.NewWorld(),
		window: window.New(window.DefaultDescriptor()),
	}
}

func (a *Application) AddStage(stage ecs.Stage) *Application {
	a.world.AddStage(stage)
	return a
}

func (a *Application) AddStageBefore(existingStage, stage ecs.Stage) *Application {
	a.world.AddStageBefore(existingStage, stage)
	return a
}

func (a *Application) AddStageAfter(existingStage, stage ecs.Stage) *Application {
	a.world.AddStageAfter(existingStage, stage)
	return a
}

func (a *Application) AddStartupSystem(system ecs.System) *Application {
	a.world.AddStartupSystem(system)
	return a
}

func (a *Application) AddShutdownSystem(system ecs.System) *Application {
	a.world.AddShutdownSystem(system)
	return a
}

func (a *Application) AddSystem(system ecs.System) *Application {
	a.world.AddSystem(system)
	return a
}

func (a *Application) AddSystemToStage(stage ecs.Stage, system ecs.System) *Application {
	a.world.AddSystemToStage(stage, system)
	return a
}

func (a *Application) Run() {
	for {
		a.world.Run()
		if !a.window.Update() {
			break
		}
	}
}
