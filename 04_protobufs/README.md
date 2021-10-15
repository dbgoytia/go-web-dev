# Protobufs

Learn from: https://developers.google.com/protocol-buffers/docs/gotutorial

## Why use protobufs - protocol buffers?

We're going to be writing an to-do application, in which you can read pending tasks, write tasks, and add comments to completed tasks and save those to a file. This requires *serialization*.

We decided to use protobufs for solving this problem because we simply need to write a *.proto* description of the data structure that you wish to store. From that, the protobuf compiler will create a class that implements automatic encoding and parsing of the protobuff data in an efficient binary format.

Generated class has already getters and setters readily available for us.


## App details

It is a command line app for managing address book files, encoded using protobufs. The command *add_person_go*, will add a new entry to the file, and *list_people_go* parses the data file and prints data to console.

## To compile this protobuf

```
make mod
make protoc
```

## Testing

```
make test
```


## Notes around protobuf installation

At the time of doing this, the official docs were followed to install protobuff. Basically, download the compiled bin from https://github.com/protocolbuffers/protobuf, in my case for Ubuntu. Then copy the files to a location availabile in your path. In my case:

```
$ which protoc
/usr/local/bin/protoc
```

```
$ ls -l /usr/local/include/
total 4
drwxr-xr-x 3 root root 4096 Oct 15 16:34 google
```

```
$ echo ${PATH}
/home/build/.vscode-server/bin/c13f1abb110fc756f9b3a6f16670df9cd9d4cf63/bin:/home/build/anaconda3/bin:/home/build/anaconda3/condabin:/home/build/.vscode-server/bin/c13f1abb110fc756f9b3a6f16670df9cd9d4cf63/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin:/home/build/go/bin:/usr/local/go/bin:/home/build/go/bin
```

## Details of my running platform

```
$ cat /etc/os-release 
NAME="Ubuntu"
VERSION="18.04.5 LTS (Bionic Beaver)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 18.04.5 LTS"
VERSION_ID="18.04"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=bionic
UBUNTU_CODENAME=bionic
```