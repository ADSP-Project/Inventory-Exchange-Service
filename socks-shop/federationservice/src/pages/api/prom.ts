import type { NextApiRequest, NextApiResponse } from 'next';

export const foovar =  function(req: NextApiRequest, status:number) {
    console.log("foovar");
    //console.log(req);
    const Prometheus = require('prom-client');
    Prometheus.register.getSingleMetric('http_request_total').labels(req.method, req.url, status).inc();
}