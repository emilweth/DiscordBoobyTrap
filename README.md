# DiscordBoobyTrap

DiscordBoobyTrap is a discord bot that ban every user posting in a given channel<br />
_Note: this is also a sandbox i never wrote go before please be kind_

## Installation

Use docker compose

```bash
cp .env.sample .env
vim .env # add your discord token / the honeypot channel ID
docker compose up -d
```

## Build

```bash
docker compose -f docker-compose.build.yml build
docker compose -f docker-compose.build.yml up
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)