// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.7.6;
pragma abicoder v2;

import "@balancer-labs/v2-vault/contracts/interfaces/IVault.sol";
import "@balancer-labs/v2-vault/contracts/interfaces/IAsset.sol";
import "../interfaces/ISwapper.sol";


contract BalancerV2Swap is ISwapper {
    struct BalancerPool {
        bytes32 poolId;
        address priceOracle;
        bool token0isAsset;
    }

    bytes4 constant public ASSETS_POOLS_SELECTOR = bytes4(keccak256("assetsPools(address)"));

    IVault public immutable vault;

    constructor (address _balancerVault) {
        vault = IVault(_balancerVault);
    }

    /// @inheritdoc ISwapper
    function swapAmountIn(
        address _tokenIn,
        address _tokenOut,
        uint256 _amount,
        address _siloOracle,
        address _siloAsset
    ) external override returns (uint256 amountOut) {
        bytes32 poolId = resolvePoolId(_siloOracle, _siloAsset);
        return _swapAmountIn(_tokenIn, _tokenOut, _amount, poolId);
    }

    /// @inheritdoc ISwapper
    function swapAmountOut(
        address _tokenIn,
        address _tokenOut,
        uint256 _amountOut,
        address _siloOracle,
        address _siloAsset
    ) external override returns (uint256 amountIn) {
        bytes32 poolId = resolvePoolId(_siloOracle, _siloAsset);
        return _swapAmountOut(_tokenIn, _tokenOut, _amountOut, poolId);
    }

    /// @inheritdoc ISwapper
    function spenderToApprove() external view override returns (address) {
        return address(vault);
    }

    function resolvePoolId(address _oracle, address _asset) public view returns (bytes32 poolId) {
        bytes memory callData = abi.encodeWithSelector(ASSETS_POOLS_SELECTOR, _asset);

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, bytes memory data) = _oracle.staticcall(callData);
        require(success, "pool for asset not set");

        BalancerPool memory pool = abi.decode(data, (BalancerPool));
        return pool.poolId;
    }

    function _swapAmountIn(
        address _tokenIn,
        address _tokenOut,
        uint256 _amountIn,
        bytes32 _poolId
    ) internal returns (uint256) {
        IVault.SingleSwap memory singleSwap = IVault.SingleSwap(
            _poolId, IVault.SwapKind.GIVEN_IN, IAsset(_tokenIn), IAsset(_tokenOut), _amountIn, ""
        );

        IVault.FundManagement memory funds = IVault.FundManagement(
            address(this), false, payable(address(this)), false
        );

        return vault.swap(singleSwap, funds, 1, block.timestamp);
    }

    function _swapAmountOut(
        address _tokenIn,
        address _tokenOut,
        uint256 _amountOut,
        bytes32 _poolId
    ) internal returns (uint256) {
        IVault.SingleSwap memory singleSwap = IVault.SingleSwap(
            _poolId, IVault.SwapKind.GIVEN_OUT, IAsset(_tokenIn), IAsset(_tokenOut), _amountOut, ""
        );

        IVault.FundManagement memory funds = IVault.FundManagement(
            address(this), false, payable(address(this)), false
        );

        return vault.swap(singleSwap, funds, type(uint256).max, block.timestamp);
    }
}
