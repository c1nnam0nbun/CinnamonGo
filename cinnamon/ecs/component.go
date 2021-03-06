package ecs

import "reflect"

type componentId int32

var componentCounter = int32(0)
var resourceCounter = int32(0)

type components struct {
	components      []componentId
	indices         map[reflect.Type]componentId
	resourceIndices map[reflect.Type]componentId
}

type componentDescriptor struct {
	id   componentId
	t    reflect.Type
	size int
}

func (c *components) initComponent(t reflect.Type) componentId {
	id, ok := c.indices[t]
	if !ok {
		id = componentId(componentCounter)
		c.indices[t] = id
		componentCounter++
	}
	c.components = append(c.components, id)
	return id
}

func (c *components) initResource(t reflect.Type) componentId {
	id, ok := c.resourceIndices[t]
	if !ok {
		id = componentId(resourceCounter)
		c.resourceIndices[t] = id
		resourceCounter++
	}
	c.components = append(c.components, id)
	return id
}
