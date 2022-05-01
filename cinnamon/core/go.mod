module github.com/c1nnam0nbun/cinnamon/core

go 1.18

replace (
	github.com/c1nnam0nbun/cinnamon/ecs => ../ecs
	github.com/c1nnam0nbun/cinnamon/window => ../window
)

require (
	github.com/c1nnam0nbun/cinnamon/ecs v0.0.0-00010101000000-000000000000
	//github.com/c1nnam0nbun/cinnamon/window v0.0.0-00010101000000-000000000000
)