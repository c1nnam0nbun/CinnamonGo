package main

import (
	"github.com/c1nnam0nbun/cinnamon/app"
	"github.com/c1nnam0nbun/cinnamon/ecs"
	"github.com/c1nnam0nbun/cinnamon/util"
)

func main() {
	app.NewApplication(app.ApplicationConfig{
		Systems: ecs.StageSystems{
			ecs.First: {
				ecs.NewSystem(testFirstSystemOne),
				ecs.NewSystem(testFirstSystemTwo),
			},
			ecs.Update: {
				ecs.NewSystem(testUpdateSystem),
			},
		},
		StartupSystems: ecs.Systems{
			ecs.NewSystem(setup),
		},
		Resources: ecs.Resources{
			ecs.NewResource[ComponentC](nil),
		},
	}).Run()
}

type ComponentA struct {
	value int
}

type ComponentB struct {
}

type ComponentC struct {
	value float64
}

func setup(world *ecs.World) {
	entity1 := world.CreateEntity()
	_ = ecs.AddComponent(&entity1, ComponentA{5})
	_ = ecs.AddComponent(&entity1, ComponentB{})
	entity2 := world.CreateEntity()
	_ = ecs.AddComponent(&entity2, ComponentA{9})
	_ = ecs.AddComponent(&entity2, ComponentB{})
	_ = ecs.AddComponent(&entity2, ComponentC{})
}

func testFirstSystemOne(world *ecs.World) {
	world.Query(util.TypeOf[ComponentA]()).ForEach(func(result ecs.QueryResult) {
		a := ecs.MatchQueryResultPtr[ComponentA](result)
		a.value++
	})
}

func testFirstSystemTwo(world *ecs.World) {
	world.Query(util.TypeOf[ComponentA]()).ForEach(func(result ecs.QueryResult) {
		a := ecs.MatchQueryResultPtr[ComponentA](result)
		a.value--
	})
}

func testUpdateSystem(world *ecs.World) {
	world.Query(util.TypeOf[ComponentA]()).ForEach(func(result ecs.QueryResult) {
		a := ecs.MatchQueryResult[ComponentA](result)
		println(a.value)
	})
}
