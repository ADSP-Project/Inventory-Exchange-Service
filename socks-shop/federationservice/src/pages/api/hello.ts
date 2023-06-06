// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import connection from '../../utils/db'

type Data = {
  name: string
}

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  if (req.method === 'GET') {
    connection.query('SELECT * FROM socks', (err, rows) => {
      if (err) {
        console.log('Error connecting to DB: ', err);
        return;
      }
      console.log('Connection established');
      res.status(200).json(rows)
    });
    res.status(200).json({ name: 'John Doe' })
  }
}
