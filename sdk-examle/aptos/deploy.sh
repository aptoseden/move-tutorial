#!/bin/bash
Profile=yekai
if [ $# -lt 3 ]; then
    echo "./deploy [codedir] [tomladdr] [PROFILE]";
    exit 0;
fi

CodeDir=$1
TomlAddr=$2
Profile=$3

echo $Profile

aptos move compile --package-dir $CodeDir --named-addresses $TomlAddr=$Profile

aptos move publish --package-dir $CodeDir --named-addresses $TomlAddr=$Profile --profile $Profile
