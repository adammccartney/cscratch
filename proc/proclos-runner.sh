#!/usr/bin/sh

TOUTER=1
TINNER=2
sleep $TOUTER &
./proclos $! $TINNER
