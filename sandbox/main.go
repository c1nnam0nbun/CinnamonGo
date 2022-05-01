package main

import (
	"fmt"
	"github.com/c1nnam0nbun/cinnamon/core"
	"github.com/c1nnam0nbun/cinnamon/ecs"
	"github.com/c1nnam0nbun/cinnamon/util"
)

func main() {
	core.NewApplication().
		AddStartupSystem(ecs.NewSystem(setup)).
		Run()
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
	world.Query(util.TypeOf[ComponentA](), util.TypeOf[ComponentB]()).ForEach(func(result ecs.QueryResult) {
		entity := ecs.MatchQueryResult[ecs.Entity](result)
		a := ecs.MatchQueryResult[ComponentA](result)
		b := ecs.MatchQueryResult[ComponentB](result)
		fmt.Println(entity, a, b)
		a.value++
	})
	world.Query(util.TypeOf[ComponentA](), util.TypeOf[ComponentB]()).ForEach(func(result ecs.QueryResult) {
		a := *ecs.MatchQueryResultPtr[ComponentA](result)
		//a := util.MatchInstancePtrA[ComponentA](components)
		//b := util.MatchInstancePtrA[ComponentB](components)
		fmt.Println(a)
	})
}
