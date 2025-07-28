# Tiltfile for TodoMVC application
# This file configures Tilt to build, deploy, and hot-reload both frontend and backend

# Load Kubernetes manifests
k8s_yaml('k8s.yaml')

# Backend (Go API)
# Build the backend Docker image and import into k3d
docker_build(
    'dev-registry:50890/todo-backend',
    context='./backend',
    dockerfile='./backend/Dockerfile',
    # only='./backend/main.go',
    # Watch for changes in Go files and rebuild
    live_update=[
        sync('./backend', '/src'),
        run('go build -o /bin/backend .', trigger=['**/*.go'])
    ]
)

# Frontend (Vue + Nginx)
# Build the frontend Docker image and import into k3d
docker_build(
    'dev-registry:50890/todo-frontend',
    context='./frontend',
    dockerfile='./frontend/Dockerfile',
    # only='./frontend/src',
    # Watch for changes in frontend files and rebuild
)

# Set up port forwarding for the backend API
k8s_resource('todo-backend', port_forwards='8080:8080')

# Set up port forwarding for the frontend
k8s_resource('todo-frontend', resource_deps=['todo-backend'], port_forwards='8081:80')
