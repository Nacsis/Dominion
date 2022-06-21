#!/bin/sh

set -e

# Define ABIGEN and SOLC default values.
ABIGEN="${ABIGEN-abigen}"
SOLC="${SOLC-solc}"

echo 'Please ensure that solc v0.7.6+ and abigen v1.10.18+ are installed.'

if ! $ABIGEN --version
then
    echo "'abigen' not found. Please add to PATH or set ABIGEN='path_to_abigen'."
    exit 1
fi

if ! $SOLC --version
then
    echo "'solc' not found. Please add to PATH or set SOLC='path_to_solc'."
    exit 1
fi

# Generate golang bindings from solidity contract
# Argument 1: solidity contract file
# Argument 2: golang contract name (used for package and file)
generate_bindings() {
    FILE=$1; PKG=$2; CONTRACT=$FILE
    echo "Generating $PKG bindings..."
    GENDIR=./generated/$PKG

    rm -r $GENDIR
    mkdir $GENDIR

    # Compile and generate binary runtime.
    $SOLC --abi --bin --bin-runtime --optimize --allow-paths contracts/vendor, $FILE.sol -o $GENDIR/
    BIN_RUNTIME=$(cat ${GENDIR}/${CONTRACT}.bin-runtime)
    OUT_FILE="$GENDIR/${CONTRACT}BinRuntime.go"
    echo "package $PKG" > $OUT_FILE
    echo >> $OUT_FILE
    echo "// ${CONTRACT}BinRuntime is the runtime part of the compiled bytecode used for deploying new contracts." >> $OUT_FILE
    echo "var ${CONTRACT}BinRuntime = \"$BIN_RUNTIME\"" >> $OUT_FILE

    # Generate bindings.
    $ABIGEN --pkg $PKG --abi $GENDIR/$FILE.abi --bin $GENDIR/$FILE.bin --out $GENDIR/$FILE.go &&\
    echo "contract compiled to $GENDIR/$FILE.go"
}

generate_bindings DominionApp dominionApp
# generate_bindings ./perun-eth-contracts/contracts/Adjudicator.sol adjudicator
# generate_bindings ./perun-eth-contracts/contracts/AssetHolderETH.sol assetHolderETH
