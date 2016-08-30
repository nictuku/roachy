# roachy
DHT supernode to learn cockroachDB

Plan:

1. run a DHT server on Scaleway.net that attracts a lot of traffic. See: https://en.wikipedia.org/wiki/Mainline_DHT
1. Record interesting bits about that traffic in a cockroachDB. Then do interesting stuff with that, like show a map of people querying us.
2. I plan to use the GetPeers logging interface to get all get_peers commands I receive and send them to the database

What I've done:

1. Copied the example code from http://github.com/nictuku/dht and started to adapt it to my needs
2. Tuned my stream settings thanks to livecodingtv :)
3. Tuned my vim settings thanks to `proton`
4. Tuned stream settings again to fix the fonts. Solution: don't downscale the video.
5. created the table
6. Made basic log insertion work!
7. Working on a geoDB integration so we can show a map of users querying the DHT
