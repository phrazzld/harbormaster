# Harbormaster

Write semantic Dockerfiles. Ship performant ones.

## Usage

Takes a Dockerfile as an argument. Outputs a new Dockerfile with some easy performance improvements made. Makes generally safe assumptions about reducing number of layers and layer size (i.e. combining subsequent RUN calls, suffixing RUN calls with `apt-get clean && rm -rf /var/lib/apt/lists/*` when `apt-get install` has been run, etc).

## License

[MIT](https://opensource.org/licenses/MIT)
