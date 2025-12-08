# Build stage
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies (use npm install since we don't have lock file yet)
RUN npm install

# Copy source code
COPY . .

# Build application
RUN npm run build

# Runtime stage
FROM node:20-alpine

WORKDIR /app

# Copy built application
COPY --from=builder /app/build ./build
COPY --from=builder /app/package*.json ./

# Install production dependencies only
RUN npm ci --production

EXPOSE 3000

ENV NODE_ENV=production

CMD ["node", "build"]
