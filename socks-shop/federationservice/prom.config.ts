const Prometheus = require('prom-client');

const httpRequestCount = new Prometheus.Counter({
  name: 'http_request_total',
  help: 'Total number of HTTP requests!',
  labelNames: ['method', 'route', 'status'],
});

const globalForPrisma = globalThis as unknown as {
  prometheus: typeof Prometheus | undefined
}

export const prometheus = globalForPrisma.prometheus ?? new Prometheus()

if (process.env.NODE_ENV !== 'production') globalForPrisma.prometheus = prometheus