# YAML Key Finder

YAML Key Finder is a simple command-line tool written in Go that allows you to search for a specific key within a YAML file and retrieve its path.

## Usage

1. **Clone the Repository**:

```bash
git clone git@github.com:vmarlier/wheredakey.git 
```

2. **Build the application**:

Build the Go application using the following command:

```bash
go build
```

This will create an executable file named `wheredakey`.

3. **Run the Application**:

Run the application by providing the target key and the path to the YAML file you want to search:

```bash
./wheredakey <target-key> <yaml-file-path>
```

Replace <target-key> with the key you're looking for and <yaml-file-path> with the path to the YAML file to search.

#### Example

Suppose you have a YAML file named example.yaml with the following content:

```yaml
toto:
    tata: ok
    titi: ok
    bobo:
         myKey: ok
```

To find the path to the key "myKey", you would run:

```bash
./wheredakey myKey example.yaml
```

The output will be:

```bash
The path to the key 'myKey' is: toto.tutu.myKey
```

### Dependencies

Go (Golang): Install the Go programming language from https://golang.org/dl/
