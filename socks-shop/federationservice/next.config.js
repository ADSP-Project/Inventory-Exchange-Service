/** @type {import('next').NextConfig} */

const nextConfig = {
  reactStrictMode: true,
  output: 'standalone',
}

const Prometheus = require('prom-client');

const httpRequestCount = new Prometheus.Counter({
  name: 'http_request_total',
  help: 'Total number of HTTP requests!',
  labelNames: ['method', 'route', 'status'],
});

module.exports = nextConfig
