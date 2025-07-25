# Tiltfile for TodoMVC application
# This file configures Tilt to build, deploy, and hot-reload both frontend and backend

# Load Kubernetes manifests
k8s_yaml('k8s.yaml')

# Backend (Go API)
# Build the backend Docker image and import into k3d
custom_build(
    'todo-backend',
    'docker build -t $EXPECTED_REF ./backend && k3d image import $EXPECTED_REF -c dev-cluster',
    deps=['./backend'],
    skips_local_docker=True,
    # Watch for changes in Go files and rebuild
    live_update=[
        sync('./backend', '/src'),
        run('go build -o /bin/backend .', trigger=['**/*.go'])
    ]
)

# Frontend (Vue + Nginx)
# Build the frontend Docker image and import into k3d
custom_build(
    'todo-frontend',
    'docker build -t $EXPECTED_REF ./frontend && k3d image import $EXPECTED_REF -c dev-cluster',
    deps=['./frontend'],
    skips_local_docker=True,
    # Watch for changes in frontend files and rebuild
    live_update=[
        # All sync steps must come first
        sync('./frontend/src', '/app/src'),
        sync('./frontend/package.json', '/app/package.json'),
        sync('./frontend/vite.config.js', '/app/vite.config.js'),
        # All run steps must come after sync steps
        run('npm run build', trigger=['src/**/*', 'package.json', 'vite.config.js']),
        run('cp -r /app/dist/* /usr/share/nginx/html/', trigger=['src/**/*', 'package.json', 'vite.config.js'])
    ]
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
