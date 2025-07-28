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
k8s_resource('todo-frontend', port_forwards='8081:80')

# Optional: Set resource dependencies
# This ensures backend starts before frontend if needed
# k8s_resource('todo-frontend', resource_deps=['todo-backend'])

# Print helpful information
print("""
🚀 TodoMVC Application Starting!

Backend API will be available at: http://localhost:8080
Frontend will be available at: http://localhost:8081

Hot reloading is enabled:
- Backend: Changes to .go files will trigger rebuild
- Frontend: Changes to src/ files will trigger rebuild

Use 'tilt up' to start and 'tilt down' to stop.
""")
