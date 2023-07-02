import type { NextApiRequest, NextApiResponse } from 'next';
import connection from '../../utils/db';

type Product = {
  id: string;
  name: string;
  description: string;
  picture: string;
  price: number;
};

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Product | { error: string }>
) {
  if (req.method === 'GET') {
    const productId = req.query.id;

    if (!productId) {
      res.status(400).json({ error: 'Product ID is required' });
      return;
    }

    connection.query(
      'SELECT * FROM `sock` WHERE sock_id = ?',
      [productId],
      function (error, results) {
        if (error) {
          console.log('Error connecting to DB:', error);
          res.status(500).json({ error: 'Internal server error' });
          return;
        }

        if (results.length === 0) {
          res.status(404).json({ error: 'Product not found' });
          return;
        }

        const product = results[0];

        // Transform the response structure
        const modifiedProduct: Product = {
          id: 'SKSH:' + product.sock_id,
          name: product.name,
          description: product.description,
          picture: product.image_url_1,
          price: product.price,
        };

        res.status(200).json(modifiedProduct);
      }
    );
  }
}
