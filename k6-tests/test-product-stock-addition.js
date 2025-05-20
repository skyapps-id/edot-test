import http from 'k6/http';
import { check } from 'k6';

export const options = {
  vus: 5,          // 5 concurrent users
  iterations: 10,  // run the default function only once
};

const url = 'http://localhost:8082/internal/warehouses/product-stock-addition';

const payload = JSON.stringify({
  products: [
    {
      product_uuid: 'cdc416b0-796c-48db-89ab-af101ceefe80',
      warehouse_uuid: '7ed81f33-cdb9-4e87-982f-cccc1e978d9e',
      quantity: 1,
    },
  ],
});

const headers = {
  'Content-Type': 'application/json',
  'Static-Token': '321',
};

export default function () {
  const res = http.post(url, payload, { headers });

  check(res, {
    'status is 200': (r) => r.status === 200,
    'status is 422': (r) => r.status === 422,
    'status is 500': (r) => r.status === 500,
  });
}
