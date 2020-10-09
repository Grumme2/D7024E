package d7024e

// The Key value store that stores the data
var KeyValueStore = make(map[string]string)

// the static number of bytes in a KademliaID
const IDLength = 20

//The maximum size of a bucket, AKA k
const bucketSize = 20
