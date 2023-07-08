import type { NextApiRequest, NextApiResponse } from 'next'

// Prometheus.collectDefaultMetrics();
export default function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  if (req.method === 'GET') {
    const Prometheus = require('prom-client');
    Prometheus.register.metrics().then((metrics:string) => {
      res.status(200).end(metrics)});
  };
};
