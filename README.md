# To-Do App Study
This is my playground for experimenting in Go, TypeScript, React, Docker, API Design, Unit Testin, etc.

## Screenshot
![To-Do App Study](https://github.com/detohm/todo-app-study/blob/main/docs/todo-app.jpeg?raw=true)

## Experiments
- Apply dependency injection approach for Go backend
- Use layer architecture for separating http(transport), service and repository layers
- Experiment with Go Unit Test per layers using mocking dependencies
- Use Go Generate with Mockgen for automatically generating the mocking dependencies
- Experiment with pprof and bench
- Use migration files with Go Migrate for the database changes
- Use Air for backend developement with hot reloading 
- Separate dockerfile for development and production build
- Use multi-stage build for production docker image
- Use ReactJS for frontend with TypeScript 
- Use atomic designn for frontend component structure
- Experiment with React Hook such as state hooks and effect hooks

## Build & Run
For development with hot reloading on both frontend and backend,
```shell
# backend hot reloading (default port 5000)
docker-compose up
# frontend hot reloading (default port 3000)
cd frontend && yarn start
```
For production build,
```shell
# build frontend
cd frontend && yarn build
# this containerised artifact have to be deployed in container orchestrations such as Kubernetes, DockerSwarm, etc.
docker build . -t todo-app-study
```
## License
MIT, Attaphong Rattanaveerachanon