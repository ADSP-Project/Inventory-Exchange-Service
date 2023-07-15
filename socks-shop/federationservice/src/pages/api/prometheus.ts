import { getMetrics } from '../../../prom.config';
import Prometheus from 'prom-client';
import type { NextApiRequest, NextApiResponse } from 'next'

type Data = {
  name: string
}

const counter = new Prometheus.Counter({
  name: 'http_request_total',
  help: 'Total number of HTTP requests!',
  labelNames: ['method', 'route', 'status'],
});

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  if (req.method === 'GET') {
    console.log('Prometheus');
    getMetrics().then((metrics: string) => {
      console.log("metrics");
      console.log(metrics);
      res.status(200).write(metrics);
      res.end();
    });
  }
}
