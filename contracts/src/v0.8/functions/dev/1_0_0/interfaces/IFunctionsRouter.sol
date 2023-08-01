// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IRouterBase} from "./IRouterBase.sol";
import {FulfillResult} from "./FulfillResultCodes.sol";

/**
 * @title Chainlink Functions Router interface.
 */
interface IFunctionsRouter is IRouterBase {
  /**
   * @notice The identifier of the route to retrieve the address of the access control contract
   * The access control contract controls which accounts can manage subscriptions
   * @return id - bytes32 id that can be passed to the "getContractById" of the Router
   */
  function getAllowListId() external pure returns (bytes32);

  /**
   * @notice Returns configuration of the Functions Router contract
   * @return adminFee Flat fee (in Juels of LINK) that will be paid to the Router owner for operation of the network
   * @return handleOracleFulfillmentSelector The function selector that is used when calling back to the Client contract
   * @return maxCallbackGasLimits List of max callback gas limits used by flag with GAS_FLAG_INDEX
   */
  function getConfig() external view returns (uint96, bytes4, uint32[] memory);

  /**
   * @notice Sends a request (encoded as data) using the provided subscriptionId
   * @param subscriptionId A unique subscription ID allocated by billing system,
   * a client can make requests from different contracts referencing the same subscription
   * @param data Encoded Chainlink Functions request data, use FunctionsClient API to encode a request
   * @param dataVersion Gas limit for the fulfillment callback
   * @param callbackGasLimit Gas limit for the fulfillment callback
   * @param donId An identifier used to determine which route to send the request along
   * @return requestId A unique request identifier
   */
  function sendRequest(
    uint64 subscriptionId,
    bytes calldata data,
    uint16 dataVersion,
    uint32 callbackGasLimit,
    bytes32 donId
  ) external returns (bytes32);

  /**
   * @notice Sends a request (encoded as data) to the proposed contracts
   * @param subscriptionId A unique subscription ID allocated by billing system,
   * a client can make requests from different contracts referencing the same subscription
   * @param data Encoded Chainlink Functions request data, use FunctionsClient API to encode a request
   * @param dataVersion Gas limit for the fulfillment callback
   * @param callbackGasLimit Gas limit for the fulfillment callback
   * @param donId An identifier used to determine which route to send the request along
   * @return requestId A unique request identifier
   */
  function sendRequestToProposed(
    uint64 subscriptionId,
    bytes calldata data,
    uint16 dataVersion,
    uint32 callbackGasLimit,
    bytes32 donId
  ) external returns (bytes32);

  /**
   * @notice Fulfill the request by:
   * - calling back the data that the Oracle returned to the client contract
   * - pay the DON for processing the request
   * @dev Only callable by the Coordinator contract that is saved in the commitment
   * @param requestId The identifier for the request
   * @param response response data from DON consensus
   * @param err error from DON consensus
   * @param juelsPerGas -
   * @param costWithoutFulfillment -
   * @param transmitter -
   * @return fulfillResult -
   * @return callbackGasCostJuels -
   */
  function fulfill(
    bytes32 requestId,
    bytes memory response,
    bytes memory err,
    uint96 juelsPerGas,
    uint96 costWithoutFulfillment,
    address transmitter
  ) external returns (FulfillResult, uint96);

  /**
   * @notice Validate requested gas limit is below the subscription max.
   * @param subscriptionId subscription ID
   * @param callbackGasLimit desired callback gas limit
   */
  function isValidCallbackGasLimit(uint64 subscriptionId, uint32 callbackGasLimit) external view;
}
