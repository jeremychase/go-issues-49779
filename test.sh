#!/bin/sh

echo "boottime 1000000000 0" > /proc/self/timens_offsets; uptime -s