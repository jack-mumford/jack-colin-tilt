# Tiltfile for TodoMVC application
# This file configures Tilt to build, deploy, and hot-reload both frontend and backend
load('ext://restart_process', 'docker_build_with_restart')

# Load Kubernetes manifests
k8s_yaml('k8s.yaml')

# Backend (Go API)
local_resource(
    'backend-compile',
    'cd backend && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/backend',
    deps=['./backend'],
    ignore=['./backend/bin']
)

docker_build_with_restart(
    'dev-registry:50890/todo-backend',
    context='./backend',
    dockerfile='./backend/Dockerfile',
    entrypoint='/app/bin/backend',
    only=[
        './bin'
    ],
    live_update=[
        sync('./backend/bin', '/app/bin'),
    ]
)

# Frontend (Vue + Nginx)
local_resource(
    'frontend-compile',
    'cd frontend && npm ci && npm run build',
    deps=['./frontend'],
    ignore=['./frontend/dist', './frontend/node_modules']
)
docker_build(
    'dev-registry:50890/todo-frontend',
    context='./frontend',
    dockerfile='./frontend/Dockerfile',
    only=[
        './dist',
        './nginx.conf'
    ],
    live_update=[
        sync('./frontend/dist', '/usr/share/nginx/html')
    ]
)

# Set up port forwarding for the backend API
k8s_resource('todo-backend', port_forwards='8080:8080')

# Set up port forwarding for the frontend
k8s_resource('todo-frontend', resource_deps=['todo-backend'], port_forwards='8081:80')
