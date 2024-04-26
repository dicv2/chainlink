// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "forge-std/Test.sol";

import {CBOREncoder} from "./CBOREncoder.sol";

// forge test --match-path src/v0.8/automation/dev/CBOREncoder.t.sol

contract SetUp is Test {
  CBOREncoder internal encoder;

  function setUp() public {
    encoder = new CBOREncoder();
  }
}

contract CBOREncode is SetUp {
  function testEncode() public {
    (bytes memory data, uint256 depth) = encoder.encode(1_000_000_000_000, 16);

    assertEq(depth, 0);
  }
}
