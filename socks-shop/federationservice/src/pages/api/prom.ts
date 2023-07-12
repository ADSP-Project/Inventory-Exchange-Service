import type { NextApiRequest, NextApiResponse} from 'next';
//const Prometheus = require('prom-client');

import { prometheus } from '../../../prom.config'


export const foovar =  function(req: NextApiRequest, status:number) {
    console.log("foovar");
    console.log(req.method);
    console.log(req.url);
    console.log(status);
    //@ts-ignore
    prometheus.register.getSingleMetric('http_request_total').labels(req.method, req.url, status).inc();
}