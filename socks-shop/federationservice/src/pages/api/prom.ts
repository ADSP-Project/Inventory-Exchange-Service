import type { NextApiRequest, NextApiResponse} from 'next';
const Prometheus = require('prom-client');


export const foovar =  function(req: NextApiRequest, status:number) {
    console.log("foovar");
    console.log(req.method);
    console.log(req.url);
    console.log(status);
    Prometheus.register.getSingleMetric('http_request_total').labels(req.method, req.url, status).inc();
}