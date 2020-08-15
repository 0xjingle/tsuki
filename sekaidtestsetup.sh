#!/bin/bash

rm -rf $HOME/.tsukid/

cd $HOME

tsukid init --chain-id=testing testing
tsukid keys add validator
tsukid add-genesis-account $(tsukid keys show validator -a) 1000000000stake,1000000000validatortoken
tsukid gentx-claim validator
tsukid start
