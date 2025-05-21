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

When('I create an order product {string} quantity {int}', async function (productUuid, qty) {
  const payload = {
    order_items: [{ product_uuid: productUuid, quantity: qty }],
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
    },
  });
});

Then('the response status should be {int}', function (statusCode) {
  expect(this.response.status).to.equal(statusCode);
});

Then('product stock is available', function () {
  expect(this.response.data.data.stock).to.greaterThan(0);
});
