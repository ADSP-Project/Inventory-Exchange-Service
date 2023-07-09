const Prometheus = require('prom-client');

console.log("next config is run!")
const httpRequestCount = new Prometheus.Counter({
  name: 'http_request_total',
  help: 'Total number of HTTP requests!',
  labelNames: ['method', 'route', 'status'],
});

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  output: 'standalone',
}

module.exports = nextConfig
