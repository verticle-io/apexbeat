# Apexbeat

Welcome to Apexbeat.

APEX extracts contextual data and metrics directly from your Java application.

It helps getting better visibility and understanding of what is happening in the software products' application layer during runtime.
Designed to accelerate application analytics, debugging and monitoring.

apexbeat makes it easy to push your extracted data to an ElasticSearch node.

![Apex Toolkit](https://verticle-io.github.io/apexbeat/vio-apex-isometric5.svg)

Visit [toolkits.verticle.io](http://toolkits.verticle.io) for more details on APEX.
	

Ensure that this folder is at the following location:
`${GOPATH}/github.com/verticle-io`

## Getting Started with Apexbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Apexbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Apexbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/verticle-io/apexbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Apexbeat run the command below. This will generate a binary
in the same directory with the name apexbeat.

```
make
```


### Run

To run Apexbeat with debugging output enabled, run:

```
./apexbeat -c apexbeat.yml -e -d "*"
```


### Test

To test Apexbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/apexbeat.template.json and etc/apexbeat.asciidoc

```
make update
```


### Cleanup

To clean  Apexbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Apexbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/verticle-io
cd ${GOPATH}/github.com/verticle-io
git clone https://github.com/verticle-io/apexbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.
