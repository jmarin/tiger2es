# tiger2es
Loads Tiger Shapefile Line Address data to Elasticsearch

## Usage
tiger2es -s <state> --host <hostname> --port <port> where:
* state: State FIPS code
* hostname (optional): IP or hostname of Elasticsearch server (localhost by default)
* port (optional): Elasticsearch http port (9200 by default)
