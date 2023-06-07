// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import connection from '../../utils/db'
import { z } from "zod";

type Data = {
  name: string
}


// table: sock
// | sock_id     | varchar(40)  | NO   | PRI | NULL    |       |
// | name        | varchar(20)  | YES  |     | NULL    |       |
// | description | varchar(200) | YES  |     | NULL    |       |
// | price       | float        | YES  |     | NULL    |       |
// | count       | int(11)      | YES  |     | NULL    |       |
// | image_url_1 | varchar(40)  | YES  |     | NULL    |       |
// | image_url_2 | varchar(40)  | YES  |     | NULL    |       |

const Socks = z.object({
    sock_id: z.string(),
    name: z.string(),
    description: z.string(),
    price: z.number(),
    count: z.number(),
    image_url_1: z.string(),
    image_url_2: z.string(),
}).array();


export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  if (req.method === 'POST') {
    const socks = Socks.parse(req.body);
    connection.query('INSERT INTO sock (sock_id, name, description, price, count, image_url_1, image_url_2) VALUES ?', [socks.map((sock) => [sock.sock_id, sock.name, sock.description, sock.price, sock.count, sock.image_url_1, sock.image_url_2])], (err, rows) => {
      if (err) {
        console.log('Error connecting to DB: ', err);
        return;
      }
      console.log('Connection established');
      res.status(200).json(rows)
    });
  }
}
