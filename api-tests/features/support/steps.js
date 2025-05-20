const { Given, When, Then, setDefaultTimeout } = require('@cucumber/cucumber');
const axios = require('axios');
const { expect } = require('chai');

setDefaultTimeout(10 * 1000); // 10 seconds

Given('I log in with id {string} and password {string}', async function (id, password) {
  const res = await axios.post('http://localhost:8080/login', { id, password });
  this.token = res.data.data.token;
});

Given('I log in and get a token', async function () {
  const res = await axios.post('http://localhost:8080/login', {
    id: '081574040777',
    password: '123',
  });
  this.token = res.data.data.token;
});

When('I request the product with UUID {string}', async function (productUuid) {
  this.response = await axios.get(`http://localhost:8081/products/${productUuid}`, {
    headers: {
      Authorization: `Bearer ${this.token}`,
    },
  });
});

When('I create an order', async function () {
  const payload = {
    order_items: [{ product_uuid: 'cdc416b0-796c-48db-89ab-af101ceefe80', quantity: 1 }],
  };

  const res = await axios.post('http://localhost:8084/orders', payload, {
    headers: {
      Authorization: `Bearer ${this.token}`,
      'Content-Type': 'application/json',
    },
  });

  this.response = res;
  this.orderUUID = res.data.data.uuid;
});

When('I fetch the order by UUID', async function () {
  this.response = await axios.get(`http://localhost:8084/orders/${this.orderUUID}`, {
    headers: {
      Authorization: `Bearer ${this.token}`,
      'Content-Type': 'application/json',
      traceparent: '00-3bdb8ae136d0bba9da57d4e603b140cf-379ace2f519c3fed-01',
    },
  });
});

Then('the response status should be {int}', function (statusCode) {
  expect(this.response.status).to.equal(statusCode);
});
