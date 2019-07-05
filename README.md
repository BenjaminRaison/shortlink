# Shortlink
## What
A go http server that redirects short links.

## Why
I needed one, and I wanted to try out go.

## How to use
1. `docker pull benjaminraison/shortlink:latest`
2. Mount a mapping file into /redirects.conf
3. Profit!

### Environment variables
* `SL_PORT`: The port to run on
* `SL_MAPPING`: Path to the mappings file

## Warnings
* Shortlink is not optimised for heavy workloads. In fact, it's not optimised at all.
* This is not production-ready. Use at your own risk.

