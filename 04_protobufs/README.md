# Protobufs

Learn from: https://developers.google.com/protocol-buffers/docs/gotutorial

## Why use protobufs - protocol buffers?

We're going to be writing an to-do application, in which you can read pending tasks, write tasks, and add comments to completed tasks and save those to a file. This requires *serialization*.

We decided to use protobufs for solving this problem because we simply need to write a *.proto* description of the data structure that you wish to store. From that, the protobuf compiler will create a class that implements automatic encoding and parsing of the protobuff data in an efficient binary format.

Generated class has already getters and setters readily available for us.


## App details

It is a command line app for managing address book files, encoded using protobufs. The command *add_person_go*, will add a new entry to the file, and *list_people_go* parses the data file and prints data to console.

