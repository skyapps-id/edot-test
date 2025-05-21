Feature: Order Product

  Scenario: Create and retrieve an order
    Given I log in and get a token
    When I request the product with UUID "cdc416b0-796c-48db-89ab-af101ceefe80"
    Then product stock is available
    When I create an order product "cdc416b0-796c-48db-89ab-af101ceefe80" quantity 1
    Then the response status should be 200
    When I fetch the order by UUID
    Then the response status should be 200
