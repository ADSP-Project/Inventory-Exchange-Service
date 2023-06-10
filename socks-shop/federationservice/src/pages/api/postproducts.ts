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


// curl -X POST -H "Content-Type: application/json" -d '[
//     {
//        "sock_id":"03fef6ac-1896-4ce8-bd69-b798f85c6e01",
//        "name":"Holy",
//        "description":"Socks fit for a Messiah. You too can experience walking in water with these special edition beauties. Each hole is lovingly proggled to leave smooth edges. The only sock approved by a higher power.",
//        "price":99.99,
//        "count":1,
//        "image_url_1":"/catalogue/images/holy_1.jpeg",
//        "image_url_2":"/catalogue/images/holy_2.jpeg"
//     }
//  ]' http://localhost:8083/api/postproducts
  

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
      console.log('Intserted socks: ', socks);
      res.status(200)
    });
  }
}
