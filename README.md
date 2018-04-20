# Status

For times when proper status monitoring is a little overkill.

## Rationale

I built this to scratch an itch with my [home network](https://fergus.london/pimping-out-your-home-network/) - I run a 4 machine Pi cluster (as a Docker Swarm), an NAS, and a few other machines. I wanted visibility on what was actually running, and if there were any interruptions.

Whilst I get some good visibility for the RPi cluster via [Portainer](https://github.com/portainer/portainer) and [Docker Swarm Visualizer](https://github.com/dockersamples/docker-swarm-visualizer), this doesn't work so well for the NAS, or a few RPi's running non-container applications. (i.e [Octoprint](https://octoprint.org/) or a DNS server.) 

## Whats it actually do?

Good question; this repeatedly `pings` a given host at a set interval - determined by the environmental variable `MONITORING_INTERVAL`. If downtime is detected, the incident is logged in a MySQL database - as well as the time at which the service was detected to be back online.

The results are accessible via an API detailed below.

### Endpoints

```
GET: /status
- Returns the status of all domains being monitored.

GET: /status/down
- Returns the status of all domains being monitored that are currently unavailable.

GET: /status/service/{domain}
- Returns the status for a specific domain.

GET: /status/service/{domain}/incidents
- Returns a list of incidents, or periods of downtime, which have occurred for the specified domain.
```

### Payloads

There are two types of payload; (a) a `check`, and (b) an `incident`.

A `check` describes the status of a given domain at a given time, and is fairly self-explanatory.

```
{
  id:             integer
  domain:         string
  last_performed: Date
  status:         bool
}
```

Whilst an `incident` describes a period of downtime - either resolved or ongoing:

```
{
  id:             integer
  check_id:       integer
  description:    string
  down_detection: Date
  up_detection:   null|Data
}
```

### But why `ping`?

Long answer: I wanted to detect whether the host itself was up, not the services running on the host. This meant doing things like a HTTP request and ensuring I get a `200` back weren't possible.. as why would DNS server respond to a HTTP request? With this in mind, I opted for ping because it's simple, I'd managed to find quite a flexible library, and it's pretty much what it's meant for.

Short answer: *my network, my rules*. I know that ping is going to work as I don't have any firewall rules preventing it (at least not internally).

### A note about the containers

The `app` container does not contain any tooling required for development: i.e `go` or `glide`. It simply mounts the executable generated via `go build` to `/app/status`, meaning **you must build the executable before running `docker-compose up`**. I will most likely take a similar approach for the `frontend` container too.

The `db` container manages persistence via mounting `docker/volumes/mysql` to `/var/lib/mysql`; for obvious reasons the contents of the mysql directory are ignored - but the `.gitkeep` folder is explicitly included via `.gitignore` to prevent having to recreate it after every clone.
