// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import connection from '../../utils/db'

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


export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  if (req.method === 'GET') {
      res.status(200)
    };

  }
