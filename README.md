#  DIRT  - Latency and Networking monitoring tool

D.I.R.T. stands for Distributed Isochronous Roundtrip Testing. It's a monitoring tool you can deploy to continuously gather latency metrics.

It provides observability metrics (min, max, mean, median and standard deviation) for : 
 - **Round Trip Time latency** (RTT): How much time a packet take to go to a remote host and then return 
 - **Send Delay**: Which is the latency to send packets to a remote host
 - **Receive Delay**: Which is the latency to receive packets from a remote host
 - **IP Packet Delay Variation** (IPDV): variation in delay (jitter) for consecutive packets for a complete round trip
 - **Send IPDV**: variation in delay (jitter)for consecutive packets to remote server
 - **Receive IPDV**: variation in delay (jitter) for consecutive packets from remote server

IPDV jitters are useful to monitor how different workloads can conflict themselves, for example with low latency small packets vs high bandwidth use cases.

## How does it work?

DIRT leverages [IRTT](https://github.com/heistp/irtt) and make it distributed. Under the hood, it uses Memberlist to create a cluster while adding filtering logic to help managing topologies.

## Clustering

You can define topologies to decide if a node should participate or not in a monitoring session. Basically, you will assigne a node into a `Group`, then define `Peers`.

`Peers` are a list of groups from which this node will accept to peer with and participate in an IRTT session.
Usually a node would want to peer with groups where it's sending or receiving traffic. This allows to avoid peering with node of the same role that don't represent a real traffic pattern. For example, two proxies are receiving traffic and connecting to backends, but they probably don't talk together this is why a node could decide to not accept peering with it's own group.

See this example from the sample config file:

```
group: backend

peers:
  - frontend
  - database
```

This node is a `backend` node, and it's not interested in checking latencies with other backend nodes. However, it's configured to accepted and establish sessions with `frontend` nodes and `database` nodes. This way, we ensure checking latencies for a real traffic pattern.

When applied to distributed clusters, you can choose to peer with node from your own group. Note that the metrics cardinality can be very high in case you apply it globally in an 100+ nodes cluster. Later enhancement would allow to enable or disable metrics to observe only a subset of metrics.


### Bootstrap node

You can define a node as a `bootstrap` node that will only manage clustering, but won't participate in latency monitoring.

### Rack awareness

In a future release, there will be the ability to elect a given node from a rack, to reach out other elected nodes from other racks. This way we reduce session cardinalities to reduce flows between machines while gathering insightful data for your topology. Usually you would want to check networking latencies between all machines among a Top of Rack Switch, while you would want to establish only 1 flow from a rack to another to avoid duplicate metrics.

## Configuration

DIRT comes with a [sample config file](config.sample.yaml). Simply copy the sample to *config.yaml*, and replace fields according to your expectations.


## Status

DIRT is considered under developpement and provided without any warranty.

Currently it integrates with [Beamium](https://github.com/clevercloud/beamium) to export Metrics. 

## Metrics

In addition to latencies metrics, DIRT exposes its own metrics with the number of nodes in the cluster. All metrics use labels to perform later aggregations and reduction (e.g. group by source / remote server). These labels are: 
- node Source
- node Destination
- rack Source
- rack Destination
- group Source
- group Destination

It will help to group by node, rack or group to identify traffic issues.

## Contributing

Contributions are welcomed in any area like:
- Code and features
- Documentation
- Logo

## Building 

Dirt is easy to build and only require having a GO runtime.

```
git clone https://github.com/StevenLeRoux/dirt.git && cd dirt
go build
```

## Get in touch

- Twitter: [@GwinizDu](https://twitter.com/GwinizDu)