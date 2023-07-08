import type { NextApiRequest, NextApiResponse } from 'next';
import connection from '../../utils/db';
import {foovar} from './prom';

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<string | { error: string }>
) {
  //foovar(req,res.statusCode);
  if (req.method === 'GET') {
    const productId = req.query.id;

    if (!productId) {
      res.status(400).json({ error: 'Product ID is required' });
      //return;
    }

    connection.query(
      'SELECT * FROM `sock` WHERE sock_id = ?',
      [productId],
      function (error, results) {
        if (error) {
          console.log('Error connecting to DB:', error);
          res.status(500).json({ error: 'Internal server error' });
          //foovar(req,res.statusCode);
          //return;
        }

        if (results.length === 0) {
          res.status(404).json({ error: 'Product not found' });
          //foovar(req,res.statusCode);
          //return;
        }

        const product = results[0];
        console.log("Current product count: " + product.count);

        if (product.count >= 1) {
            const updateCountQuery = 'UPDATE sock SET count = count - 1 WHERE sock_id = ?';

            connection.query(updateCountQuery, [productId], (err, result) => {
            if (err) {
                console.log('Error updating count: ', err);
                return;
            }
            console.log(result);
            console.log('Count updated successfully');
            res.status(200).json("Success")
            });
        } else {
            res.status(400).json("Product not available")
        }
        //foovar(req,res.statusCode);
      }
    );

    foovar(req,res.statusCode);
  }
}
