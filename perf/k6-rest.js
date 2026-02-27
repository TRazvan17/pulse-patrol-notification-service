import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 50,
  duration: '30s',
  thresholds: {
    http_req_failed: ['rate<0.01'],
    http_req_duration: ['p(95)<300'],
  },
};

export default function () {
  const url = 'http://localhost:8080/notifications';
  const payload = JSON.stringify({ to: 'lector', message: 'hello from k6' });

  const params = { headers: { 'Content-Type': 'application/json' } };

  const res = http.post(url, payload, params);
  check(res, { 'status is 200': (r) => r.status === 200 });

  sleep(0.1);
}