# Shadow Link

Shadow Link is a bot that creates a collaborative playlist of music and video links shared in your Discord server.

- Add Shadow Link to your Discord server
- Paste a YouTube link in a channel
- See 🔗 reaction from Shadow Link if it saved the link
- Add ❌ reaction to delete the like

PRIVACY WARNING: Shadow Link listens to all messages in your server and exfiltrates some content. Be careful.

## Roadmap

- [x] Extract and save links from messages to database
- [x] Extract opengraph metadata from web pages
- [ ] Delete links by reaction
- [ ] Send all bot activity to a mod channel
- [ ] Generate public feeds
  - [ ] RSS feed
  - [ ] Spotify playlist
  - [x] Web page
  - [ ] YouTube playlist

## Deploy

```bash
flyctl apps create shadowlink
make deploy
```
