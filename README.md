# TodoMVC Monorepo (Vue 3 + Go)

A minimal **TodoMVC** application built as two micro-services inside a single repository:

* **frontend/** – Vue 3 + Vite single-page app
* **backend/**  – Go 1.22 HTTP JSON API

## Prerequisites

* Go 1.22+
* Node 18+ / npm 9+

## Quick start

```bash
# 1. install JS deps
cd frontend
npm install

# 2. in a separate terminal – run the backend
cd ../backend
go run .

# 3. back to the frontend – start dev server (will open http://localhost:5173)
cd ../frontend
npm run dev -- --host
```

Now open `http://localhost:5173` and start adding todos. The SPA talks to the Go API at `http://localhost:8080`.

## Environment configuration

Vite uses environment files to embed the backend URL at build time. Two files live in `frontend/` (and are **git-ignored**):

| File | Used by | Example value |
|------|---------|---------------|
| `.env.development` | `npm run dev` | `VITE_API_BASE_URL=http://localhost:8080` |
| `.env.production`  | `npm run build` / Docker / k8s | `VITE_API_BASE_URL=http://todo-backend:8080` |

Only variables prefixed with `VITE_` are exposed to the client bundle. Feel free to create additional files such as `.env.staging` and pass them inline:

```bash
VITE_API_BASE_URL=https://staging.example.com npm run build
```

---

## Project structure

```
├── backend           # Go service (./main.go, go.mod)
│   └── ...
├── frontend          # Vue 3 SPA (Vite)
│   ├── index.html
│   └── src/
│       ├── App.vue
│       └── main.js
└── .gitignore        # shared ignores
```

## API

| Method | Endpoint        | Description               |
| ------ | --------------- | ------------------------- |
| GET    | /todos          | List todos                |
| POST   | /todos          | Create `{ text }`         |
| PUT    | /todos/{id}     | Replace a todo            |
| DELETE | /todos/{id}     | Delete a todo             |

All responses are JSON. No authentication; everything is in-memory for demo purposes.

## Building for production

Frontend build:

```bash
cd frontend
npm run build   # outputs to frontend/dist
```

You can then serve `dist/` with any static server (e.g., Nginx) while running the Go service.

---
MIT © 2025
