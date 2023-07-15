/** @type {import('next').NextConfig} */

const nextConfig = {
  reactStrictMode: true,
  output: 'standalone',
}

// const Prometheus = require('prom-client');
// //const register = new Prometheus.Registry();

// const httpRequestCount = new Prometheus.Counter({
//   name: 'http_request_total',
//   help: 'Total number of HTTP requests!',
//   labelNames: ['method', 'route', 'status'],
// });

//register.registerMetric(http_request_counter);

module.exports = nextConfig
