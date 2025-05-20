import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 10,        // 10 concurrent users
  duration: '10s' // max duration (can be a bit longer to ensure all requests finish)
};

export function setup() {
  const loginPayload = JSON.stringify({
    id: '081574040777',
    password: '123',
  });

  const loginHeaders = { 'Content-Type': 'application/json' };

  const loginRes = http.post('http://localhost:8080/login', loginPayload, {
    headers: loginHeaders,
  });

  check(loginRes, {
    'login status is 200': (res) => res.status === 200,
    'login token present': (res) => !!res.json('data.token'),
  });

  const token = loginRes.json('data.token');
  return { token };
}

export default function (data) {
  const token = data.token;

  const productUrl = 'http://localhost:8081/products/cdc416b0-796c-48db-89ab-af101ceefe80';

  // Each VU makes 10 requests sequentially
  for (let i = 0; i < 10; i++) {
    const productRes = http.get(productUrl, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });

    check(productRes, {
      'product status is 200': (res) => res.status === 200,
      'response time < 500ms': (res) => res.timings.duration < 500,
    });

    sleep(1); // Optional delay between requests per user
  }
}
