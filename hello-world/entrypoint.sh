#!/bin/sh
sysctl kernel.randomize_va_space=0
exec /bin/bash -l
