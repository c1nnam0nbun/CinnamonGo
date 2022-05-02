package app

import (
	"github.com/c1nnam0nbun/cinnamon/ecs"
	"github.com/c1nnam0nbun/cinnamon/window"
)

type Application struct {
	world  *ecs.World
	window window.Window
}

type Plugin interface {
	Build(application *Application)
}

type Plugins []Plugin

type ApplicationConfig struct {
	Plugins         Plugins
	Systems         ecs.StageSystems
	StartupSystems  ecs.Systems
	ShutdownSystems ecs.Systems
	CustomStages    ecs.Systems
	Resources       ecs.Resources
}

func NewApplication(cfg ApplicationConfig) *Application {
	a := Application{
		world:  ecs.NewWorld(),
		window: window.New(window.DefaultDescriptor()),
	}
	a.setup(cfg)
	for _, plugin := range cfg.Plugins {
		plugin.Build(&a)
	}
	return &a
}

func (a *Application) UpdateConfig(cfg ApplicationConfig) {
	a.setup(cfg)
}

func (a *Application) setup(cfg ApplicationConfig) {
	for k, v := range cfg.Systems {
		for _, s := range v {
			a.addSystemToStage(k, s)
		}
	}
	for _, s := range cfg.StartupSystems {
		a.addStartupSystem(s)
	}
	for _, s := range cfg.ShutdownSystems {
		a.addShutdownSystem(s)
	}
	for _, r := range cfg.Resources {
		a.initResource(r)
	}
}

func (a *Application) initResource(res ecs.Resource) {
	a.world.InitResource(res)
}

func (a *Application) addStage(stage ecs.Stage) *Application {
	a.world.AddStage(stage)
	return a
}

func (a *Application) addStageBefore(existingStage, stage ecs.Stage) *Application {
	a.world.AddStageBefore(existingStage, stage)
	return a
}

func (a *Application) addStageAfter(existingStage, stage ecs.Stage) *Application {
	a.world.AddStageAfter(existingStage, stage)
	return a
}

func (a *Application) addStartupSystem(system ecs.System) *Application {
	a.world.AddStartupSystem(system)
	return a
}

func (a *Application) addShutdownSystem(system ecs.System) *Application {
	a.world.AddShutdownSystem(system)
	return a
}

func (a *Application) addSystem(system ecs.System) *Application {
	a.world.AddSystem(system)
	return a
}

func (a *Application) addSystemToStage(stage ecs.Stage, system ecs.System) *Application {
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
