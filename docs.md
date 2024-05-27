
forge install OpenZeppelin/openzeppelin-contracts@v4.9.1 --no-commit

abigen --abi solc_abis/ISilo.abi --pkg Silo --out "bindings/ISilo/ISilo.go" 

solc contracts/BaseSilo.sol --base-path . --include-path lib/