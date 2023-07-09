import type { NextApiRequest, NextApiResponse } from 'next'

type Data = {
  name: string
}

// Prometheus.collectDefaultMetrics();
export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  if (req.method === 'GET') {
    console.log('Prometheus');
    const Prometheus = require('prom-client');
    Prometheus.register.metrics().then((metrics:string) => {
      console.log("metrics");
      console.log(metrics);
      res.status(200).write(metrics);
      res.end();
    });
      //resolve();
  };
};
