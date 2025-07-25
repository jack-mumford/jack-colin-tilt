# ---------- Builder Stage ----------
# Use a lightweight Node image to compile the Vue app
FROM node:20-alpine AS builder

# Create app directory inside the container
WORKDIR /app

# Copy dependency manifests first for better layer caching
COPY package*.json ./

# Install only production-needed deps for smaller final image, but keep dev deps during build
RUN npm ci

# Copy the rest of the source tree
COPY . .

# Build the static files into /app/dist
RUN npm run build

# ---------- Runtime Stage ----------
FROM nginx:alpine

# Replace default Nginx site config with our SPA-friendly one
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy compiled assets from builder
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

# Start Nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]
