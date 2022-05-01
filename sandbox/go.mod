module github.com/c1nnam0nbun/sandbox

go 1.18

replace (
	github.com/c1nnam0nbun/cinnamon/core => ../cinnamon/core
	github.com/c1nnam0nbun/cinnamon/ecs => ../cinnamon/ecs
	github.com/c1nnam0nbun/cinnamon/util => ../cinnamon/util
	github.com/c1nnam0nbun/cinnamon/window => ../cinnamon/window
)

require (
	github.com/c1nnam0nbun/cinnamon/core v0.0.0-20201127153200-e1b16c1ebc08
	github.com/c1nnam0nbun/cinnamon/ecs v0.0.0-20201127153200-e1b16c1ebc08
	github.com/c1nnam0nbun/cinnamon/util v0.0.0-20201127153200-e1b16c1ebc08
)

require (
	github.com/c1nnam0nbun/cinnamon/window v0.0.0-20201127153200-e1b16c1ebc08 // indirect
	github.com/cnf/structhash v0.0.0-20201127153200-e1b16c1ebc08 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220320163800-277f93cfa958 // indirect
	github.com/viant/xunsafe v0.8.0 // indirect
	golang.org/x/exp v0.0.0-20220426173459-3bcf042a4bf5 // indirect
)