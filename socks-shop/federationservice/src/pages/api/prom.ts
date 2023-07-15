import Prometheus from 'prom-client';
import type { NextApiRequest, NextApiResponse } from 'next'


const counter = new Prometheus.Counter({
  name: 'http_request_total',
  help: 'Total number of HTTP requests!',
  labelNames: ['method', 'route', 'status'],
});

export const foovar = function(req: NextApiRequest, status: number) {
  console.log("foovar");
  console.log(req.method);
  console.log(req.url);
  console.log(status);

  const labels = {
    method: req.method,
    route: req.url,
    status: status.toString(),
  };

  counter.inc(labels);
}
