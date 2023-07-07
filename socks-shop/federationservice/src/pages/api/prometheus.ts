import type { NextApiRequest, NextApiResponse } from 'next'

type Data = {
  name: string
}

// Prometheus.collectDefaultMetrics();
export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  const Prometheus = require('prom-client');
  Prometheus.register.metrics().then((metrics:string) => {
    res.status(200).end(metrics)
  })
}
