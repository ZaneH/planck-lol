docker_build('planck-lol-service', '.')

k8s_yaml('k8s/deployment.yaml')

# Forward service to localhost:8888
k8s_resource('planck-lol-service', port_forwards=8888)
