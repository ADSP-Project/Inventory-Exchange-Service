import type { NextApiRequest, NextApiResponse } from 'next'
const Prometheus = require('prom-client');

type Data = {
  name: string
}

// Prometheus.collectDefaultMetrics();
export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  if (req.method === 'GET') {
    console.log('Prometheus');
    try {
      const metrics = await Prometheus.register.metrics();
    //Prometheus.register.metrics().then((metrics:string) => {
      console.log("metrics");
      console.log(metrics);
      res.status(200).end(metrics);
      //res.send();
    //});
    }
    catch (error) {
      console.log(error);
      res.status(400).end("Error");}
  };
};
