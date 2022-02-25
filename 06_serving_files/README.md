# Serving files

A section of the training explaining the approaches to create a FileServer, along with some interesting topics to move forward after finishing with the three exercises.

A file server is a computer responsible for the storage and management of data files so that other computers on the same network can access the files. It enables users to share information over a network without having to physically transfer files.

The file server takes on the computer or server role to store and make available data blobs to clients, serving as a central location to store and share files for a network. They can be limited to a single local area network (LAN) or can be open to the internet.

File servers make storing, securing and sharing files in an organization simpler. File servers are a common target for hackers and ransomware, so particular attention must be given to securing them against attacks.


## Common features found in file servers

* File servers typically include additional features to enable multiple users to access them simultaneously.
* Permission management is used to set who can access which files and who has rights to edit or delete the files.
* File locking stops multiple users from editing the same file at the same time.
* Conflict resolution maintains data integrity in the event of files being overwritten.
* A distributed file system can make the data redundant and highly available by copying it to multiple servers at different locations.


## A note about this implementation

Of course this very simple implenetation doesn't have all of the features of a file server implemented, but rather is meant to be an academic way of learning and understanding it's most basic implementation.