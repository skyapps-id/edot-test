const { Given, When, Then } = require('@cucumber/cucumber');
const axios = require('axios');
const { expect } = require('chai');

let token = '';
let response;

Given('I log in with id {string} and password {string}', async function (id, password) {
  const res = await axios.post('http://localhost:8080/login', {
    id,
    password,
  });

  expect(res.status).to.equal(200);
  token = res.data.data.token;
});

When('I request the product with UUID {string}', async function (productUuid) {
  response = await axios.get(`http://localhost:8081/products/${productUuid}`, {
    headers: {
      Authorization: `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
  });
});

Then('the response status should be {int}', function (statusCode) {
  expect(response.status).to.equal(statusCode);
});
