#! /bin/bash

for file in deployments/arbitrum/*.json; do
    echo "Generating bindings for $file"
    file_name="$(basename $file .json)"
    mkdir "bindings/$file_name"
    (cat $file | jq -r ".address") > "bindings/$file_name/address.txt"
    (cat $file | jq -r ".abi") | abigen --pkg "$file_name" --abi /dev/stdin --out "bindings/$file_name/$file_name.go"
done