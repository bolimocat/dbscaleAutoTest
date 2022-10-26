#!/bin/bash
echo 'ifdown ...'
sleep 10
ifdown enp0s3
sleep 60
ifup enp0s3
sleep 10
