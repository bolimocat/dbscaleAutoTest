#!/bin/bash
echo 'ifdown ...'
sleep 3
ifdown enp0s3
sleep 30
ifup enp0s3
sleep 3
