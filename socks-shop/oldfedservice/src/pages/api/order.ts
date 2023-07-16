import type { NextApiRequest, NextApiResponse } from 'next'

export interface ExternalMoney {
    CurrencyCode: string;
    Units: number;
    Nanos: number;
  }
  
  export interface ExternalAddress {
    streetAddress: string;
    city: string;
    state: string;
    country: string;
    zipCode: number;
  }
  
  export interface ExternalOrderItem {
    item: string;
    cost: ExternalMoney;
  }
  
  export interface ExternalOrderData {
    orderId: string;
    shippingTrackingId: string;
    shippingCost: ExternalMoney;
    shippingAddress: ExternalAddress;
    items: ExternalOrderItem[];
  }
  

export default function handler(
    req: NextApiRequest,
    res: NextApiResponse
  ) {
    if (req.method === 'POST') {
        try {
        const postData: ExternalOrderData = req.body;
        res.status(200).json({ message: 'Order processed successfully' });
        console.log(postData);
        } catch (error) {
        // Return an error response
        res.status(500).json({ error: 'Internal server error' });
        }
    }
};