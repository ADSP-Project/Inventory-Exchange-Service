// const Prometheus = require('prom-client');
// const register = new Prometheus.Registry();

// const httpRequestCount = new Prometheus.Counter({
//   name: 'http_request_total',
//   help: 'Total number of HTTP requests!',
//   labelNames: ['method', 'route', 'status'],
// });

// register.registerMetric(httpRequestCount);

// const globalForProm = globalThis as unknown as {
//   prometheus: typeof Prometheus | undefined
// }

// export const prometheus = globalForProm.prometheus ?? new Prometheus()

// globalForProm.prometheus = prometheus

import Prometheus from 'prom-client';

const counter = new Prometheus.Counter({
  name: 'http_request_total',
  help: 'Total number of HTTP requests!',
  labelNames: ['method', 'route', 'status'],
});

export function publishMetrics(labelNames) {
  counter.inc({
    method: labelNames.method,
    route: labelNames.route,
    status: labelNames.status,
  });
}

export async function getMetrics() {
  const registry = new Prometheus.Registry();
  registry.registerMetric(counter);
  return await registry.metrics();
}