load('ext://restart_process', 'docker_build_with_restart')

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/server ./cmd/server'
local_resource('planck-lol-compile', compile_cmd, deps=['./internal'])

docker_build_with_restart(
    'planck-lol-service:dev',
    '.',
    only=[
        './build'
    ],
    entrypoint=['/app/build/server'],
    live_update=[
        sync('./build', '/app/build'),
    ]
)

k8s_yaml('k8s/app/deployment.yaml')
k8s_resource('planck-lol-service', port_forwards='8888:8080')

local('helm upgrade --install redis bitnami/redis -f k8s/redis/values.yaml')
local('kubectl apply -f k8s/postgres/citus.yaml')
