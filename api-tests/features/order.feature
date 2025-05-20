Feature: Order Product

  Scenario: Create and retrieve an order
    Given I log in and get a token
    When I create an order
    Then the response status should be 200
    When I fetch the order by UUID
    Then the response status should be 200
