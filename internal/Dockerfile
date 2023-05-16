# Go build
FROM golang:1.20 as go-build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./internal/server

# Frontend build
FROM node:18-alpine AS frontend-build

WORKDIR /app

RUN apk add --no-cache libc6-compat
COPY . .

WORKDIR /app/internal/frontend

RUN npm install

ENV NEXT_TELEMETRY_DISABLED 1

RUN npm run build


# Runner
FROM node:18-alpine AS runner

WORKDIR /app

ENV NODE_ENV production

ENV NEXT_TELEMETRY_DISABLED 1

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=frontend-build /app/internal/frontend/public ./public

COPY --from=frontend-build --chown=nextjs:nodejs /app/internal/frontend/.next/standalone ./
COPY --from=frontend-build --chown=nextjs:nodejs /app/internal/frontend/.next/static ./.next/static

COPY --from=go-build /app/app /app/server

COPY internal/run.sh .

USER nextjs

CMD ["/bin/sh", "run.sh"]