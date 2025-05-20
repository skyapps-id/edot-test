Feature: Login and Get Product

  Scenario: Successful login and product fetch
    Given I log in with id "081574040777" and password "123"
    When I request the product with UUID "cdc416b0-796c-48db-89ab-af101ceefe80"
    Then the response status should be 200
